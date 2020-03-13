from flask import Flask, request, jsonify
from keras.models import load_model
from bson.objectid import ObjectId
from sklearn import preprocessing
import pandas as pd
import pymongo

header_cate = [
    "item_brand", "item_name", "user_gender", "store_city"
]
header_cont=[
    "item_price", "item_salecount", "item_score",  "store_level", "item_timestamp", "store_timestamp"
]
header_cont_user=[
    "user_age", "user_historysum"
]

def get_enc():
    myclient = pymongo.MongoClient("mongodb://182.92.196.182:27017/")
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


def get_data(user_id):
    # 链接数据库
    myclient = pymongo.MongoClient("mongodb://182.92.196.182:27017/")
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
    enc = get_enc()
    arr = enc.transform(df_cate).toarray()
    df_cate = pd.DataFrame(arr)
    # df_cate = pd.get_dummies(df_cate)
    # print(df_cate)
    frame = [df_cont, df_cont_user, df_cate]
    X = pd.concat(frame, axis=1)
    return X, df_item


model = load_model("C://Users/Ryanw/go/src/github.com/somewhere/algo/src-go/deep_wide_model.h5")
X, df_item=get_data("5df9e1fe91560048ad6fb730")
predicted = model.predict([X, X])

app = Flask(__name__)
@app.route('/test', methods=['GET','POST'])
def index():
    if request.method == 'GET':
        return "hello"
    else:
        data = request.get_json()
        user_id = data["user_id"]
        X, df_item=get_data(user_id)
        # print("即将进行运算的user_id:",user_id)
        # print("开始计时：",time.time())
        total = []
        # X, df_item=get_data(user_id)
        # print("加载数据：",time.time()-begin)
        # unknown = np.array([[1.0,1.0,0.0,1.0,0.0,0.0,0.0,0.0,1.0,1,1,0,1,0,1,0,1,0 ]])
        # print(model)
        # print("加载模型：",time.time()-begin)       
        predicted = model.predict([X, X])
        # print("模型预测：",time.time()-begin)
        i = 0
        for _, row in df_item.iterrows():
            ans = {}
            ans['item_id'] = str(row['item_id'])
            ans['score']=float(predicted[i][0])
            total.append(ans)
            i+=1
        # print(ans)
        return jsonify({"msg":total})


if __name__ == '__main__':
    app.debug=True
    app.run()