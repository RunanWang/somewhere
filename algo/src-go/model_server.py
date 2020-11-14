from flask import Flask, request, jsonify
from tensorflow.keras.models import load_model
from bson.objectid import ObjectId
from sklearn import preprocessing
import pandas as pd
import pymongo
import threading
import train_deep_wide

header_cate = [
    "item_brand", "item_name", "user_gender", "store_city"
]
header_cont=[
    "item_price", "item_salecount", "item_score",  "store_level", "item_timestamp", "store_timestamp"
]
header_cont_user=[
    "user_age", "user_historysum"
]

basic_model = load_model("C://Users/Ryanw/go/src/github.com/somewhere/algo/src-go/deep_wide_model_new.h5")
model = basic_model


def get_enc():
    myclient = pymongo.MongoClient("mongodb://127.0.0.1:27017/")
    mydb = myclient["kit"]
    rec_col = mydb["records"]
    item_col = mydb["items"]
    user_col = mydb["users"]
    store_col = mydb["stores"]
    records = rec_col.find({},{"_id":0, "query":0})
    items = item_col.find()
    users = user_col.find()
    stores = store_col.find()

    # 将cursor转换成dataframe格式
    df_rec = pd.DataFrame(list(records))
    df_item = pd.DataFrame(list(items))
    df_user = pd.DataFrame(list(users))
    df_store = pd.DataFrame(list(stores))

    # 按照user-id对record进行排序
    df_rec = df_rec.sort_values(by="user_id", ascending=False)

    # 对列名进行重命名（将_id转换为具体的），并设置索引
    df_item = df_item.rename(columns={'_id':'item_id'})
    df_item.set_index(['item_id'],inplace=True)
    df_user = df_user.rename(columns={'_id':'user_id'})
    df_user.set_index(['user_id'],inplace=True)
    df_store = df_store.rename(columns={'_id':'store_id','timestamp':'store_timestamp'})
    df_store.set_index(['store_id'],inplace=True)

    # 根据item_id链接record和item
    df_result = df_rec.set_index(['item_id'])
    df_result = pd.concat([df_result, df_item], axis = 1, join='inner')
    df_result = df_result.reset_index()

    # 根据user_id链接record和user
    df_result.set_index(['user_id'],inplace=True)
    df_result = pd.concat([df_result, df_user], axis = 1, join='inner')
    df_result = df_result.reset_index()

    # 根据store_id链接record和store
    df_result.set_index(['store_id'],inplace=True)
    df_result = pd.concat([df_result, df_store], axis = 1, join='inner')
    df_result = df_result.reset_index()

    # 对数据按列进行分类，不同的列采用不同的处理方法
    df_cate = df_result[header_cate]

    # 对cate类型进行独热编码
    enc = preprocessing.OneHotEncoder(handle_unknown='ignore')
    enc.fit(df_cate)
    return enc

enc = get_enc()

def get_data(user_id):
    # 链接数据库
    myclient = pymongo.MongoClient("mongodb://127.0.0.1:27017/")
    mydb = myclient["kit"]
    item_col = mydb["items"]
    user_col = mydb["users"]
    store_col = mydb["stores"]
    # 扫数据
    items = item_col.find()
    users = user_col.find({"_id":ObjectId(user_id)})
    stores = store_col.find()
    # 转换成dataframe，并重命名列名
    df_item = pd.DataFrame(list(items))
    df_item = df_item.rename(columns={'_id':'item_id'})
    df_user = pd.DataFrame(list(users))
    df_user = df_user.rename(columns={'_id':'user_id'})
    # print(df_user)
    df_store = pd.DataFrame(list(stores))
    df_store = df_store.rename(columns={'_id':'store_id','timestamp':'store_timestamp'})
    df_store.set_index(['store_id'],inplace=True)
    # 组合user和item
    df_result = df_item
    df_user_col_name = df_user.columns.values.tolist()
    for col_name in df_user_col_name:
        fulfil = []
        for i in range(df_result.shape[0]):
            fulfil.append(df_user[col_name].values[0])
        df_result[col_name] = fulfil
    # 组合store信息
    df_result.set_index(['store_id'],inplace=True)
    df_result = pd.concat([df_result, df_store], axis = 1, join='inner')
    df_result = df_result.reset_index()
    # print(df_result)
    df_cont = df_result[header_cont]
    df_cont_user = df_result[header_cont_user]
    df_cate = df_result[header_cate]
    scaler = preprocessing.MinMaxScaler(feature_range=(0, 1))
    scaled = scaler.fit_transform(df_cont)
    df_cont = pd.DataFrame(scaled)
    # enc = get_enc()
    arr = enc.transform(df_cate).toarray()
    df_cate = pd.DataFrame(arr)
    # df_cate = pd.get_dummies(df_cate)
    # print(df_cate)
    # frame = [df_cont, df_cont_user, df_cate]
    frame = [df_cont,df_cont_user,df_cate]
    X = pd.concat(frame, axis=1)
    return X, df_item


def change_model(app):
    # 等待所有的rec请求都不使用旧模型
    app.using_old_model.acquire()
    app.using_old_model.wait()
    # 更换模型
    model = app.model
    print("模型已更换")
    # 重置更新状态为False
    app.model_lock.acquire()
    app.exist_new_model = False
    app.model_lock.release()
    app.using_old_model.release()
    # 重置训练模型状态为False，可以进行下一次训练
    app.lock.acquire()
    app.is_training = False
    app.lock.release()


def getItemCFMatrix():
    myclient = pymongo.MongoClient("mongodb://127.0.0.1:27017/")
    mydb = myclient["kit"]
    item_col = mydb["items"]
    rec_col = mydb["records"]

    items = item_col.find()
    df_item = pd.DataFrame(list(items))
    df_item = df_item.rename(columns={'_id': 'item_id'})
    IDtoNum = {}
    IDtoMap = {}
    for idx, item in df_item.iterrows():
        for idxx, itemi in df_item.iterrows():
            IDtoNum[itemi['item_id']] = 0
        IDtoMap[item['item_id']] = IDtoNum
        IDtoNum = {}
    # print(IDtoMap)

    records = rec_col.find()
    df_rec = pd.DataFrame(list(records))
    # print(df_rec)
    df_rec.set_index(['user_id'], inplace=True)
    df_rec = df_rec.sort_index()
    df_rec = df_rec.reset_index()
    user_id = ""
    item_list = []
    for idx, item in df_rec.iterrows():
        if user_id == "":
            user_id = item['user_id']
            item_list = []
        elif user_id != item['user_id']:
            user_id = item['user_id']
            for itemi in item_list:
                for itemj in item_list:
                    IDtoMap[itemi][itemj] = IDtoMap[itemi][itemj] + 1
                    IDtoMap[itemj][itemi] = IDtoMap[itemj][itemi] + 1
            item_list = []
        if item['is_trade'] == 1:
            item_list.append(item['item_id'])

    for itemi in item_list:
        for itemj in item_list:
            IDtoMap[itemi][itemj] = IDtoMap[itemi][itemj] + 1
            IDtoMap[itemj][itemi] = IDtoMap[itemj][itemi] + 1

    return IDtoMap


def getUserScore(user_id, m):
    m = getItemCFMatrix()
    myclient = pymongo.MongoClient("mongodb://127.0.0.1:27017/")
    mydb = myclient["kit"]
    rec_col = mydb["records"]
    item_col = mydb["items"]

    records = rec_col.find({"user_id": ObjectId(user_id)})
    df_rec = pd.DataFrame(list(records))

    items = item_col.find()
    df_item = pd.DataFrame(list(items))

    idToFreq = {}
    idToScore = {}
    total = 0
    for _, item in df_item.iterrows():
        idToFreq[item['_id']] = 0
        idToScore[item['_id']] = 0

    for _, item in df_rec.iterrows():
        if item['is_trade'] == 1:
            idToFreq[item['item_id']] = idToFreq[item['item_id']] + 1
            total = total + 1

    for k in idToFreq.keys():
        if total != 0:
            idToFreq[k] = idToFreq[k]/total
    
    for k in idToScore.keys():
        for kk in idToFreq.keys():
            idToScore[k] = idToScore[k] + idToFreq[kk] * m[kk][k]

    
    # print(idToFreq)
    # print(idToScore)

    total = 0
    ans = {}
    for k in idToScore.keys():
        total = total + idToScore[k]
    for k in idToScore.keys():
        if total!=0:
            ans[k.__str__()] = idToScore[k]/total
        else:
            ans[k.__str__()] = idToScore[k]

    return ans


app = Flask(__name__)
app.lock = threading.Lock()
app.model_lock = threading.Lock()
app.is_training = False
app.using_old_model = threading.Condition()
app.using_model_num = 0
app.exist_new_model = False
app.model = model
app.matrix = getItemCFMatrix()

@app.route('/test', methods=['GET','POST'])
def index():
    if request.method == 'GET':
        # 用于获得目前是否有模型等待进行更新
        app.model_lock.acquire()
        num = app.using_model_num
        status = app.exist_new_model
        app.model_lock.release()
        if status:
            return jsonify({"code":3,"status":"new model waiting to refresh","num":num})
        else:
            return jsonify({"code":2,"status":"no new model"})
    else:
        data = request.get_json()
        user_id = data["user_id"]
        X, df_item=get_data(user_id)
        total = []
        app.model_lock.acquire()
        # 存在新模型等待更新，用basic模型进行预测，并在using=0时释放信号
        if app.exist_new_model:
            if app.using_model_num == 0:
                app.using_old_model.acquire()
                app.using_old_model.notify()
                app.using_old_model.release()
            app.model_lock.release()
            predicted = basic_model.predict([X, X])
        # 不存在新模型等待更新，用model直接预测，并对usingnum进行操作
        else:
            app.using_model_num = app.using_model_num + 1
            app.model_lock.release()
            predicted = model.predict([X, X])   
            app.model_lock.acquire()
            app.using_model_num = app.using_model_num - 1
            app.model_lock.release()
        i = 0
        CFscore = getUserScore(user_id, app.matrix)
        print(CFscore)
        for _, row in df_item.iterrows():
            ans = {}
            ans['item_id'] = str(row['item_id'])
            ans['score']=float(predicted[i][0]) + CFscore[str(row['item_id'])]
            total.append(ans)
            i+=1
        return jsonify({"msg":total})


@app.route('/status', methods=['GET'])
def index2():
    # 用于获得目前是否有模型正在进行训练
    app.lock.acquire()
    status = app.is_training
    app.lock.release()
    if status:
        return jsonify({"code":1, "status":"training"})
    else:
        return jsonify({"code":0,"status":"serving"})


@app.route('/train', methods=['POST'])
def index3():
    # 进入训练，先看是否有正在训练的模型
    app.lock.acquire()
    # 若存在正在训练的模型，直接resp返回1531
    if app.is_training:
        app.lock.release()
        return jsonify({"err_code":1503, "err_msg":"Another model is training."})
    # 否则将训练状态置成True
    else:
        app.is_training = True
        app.lock.release()
    # 然后进行训练，启动两个线程
    try:
        # t1负责将新模型放入app.model中
        t1 = threading.Thread(target=train_deep_wide.build_serve, args=("./deep_wide.h5",app,))
        t1.start()
        # t2负责更新新模型
        t2 = threading.Thread(target=change_model, args=(app,))
        t2.start()
    except:
        # 特殊情况是线程不能启动，报错1003
        return jsonify({"err_code":1003, "err_msg":"Something wrong."})
    # 正常返回0
    return jsonify({"err_code":0, "err_msg":"OK, training."})





if __name__ == '__main__':
    app.debug=True
    app.run()