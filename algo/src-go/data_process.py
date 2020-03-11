import pymongo
import pandas as pd
from sklearn import preprocessing

# 一些列名，按照列名对不同变量进行不同的初始化处理。
header_id = [
    "store_id", "user_id", "item_id"
]
header_cate = [
    "item_brand", "item_name",  "user_gender", "store_city"
]
header_cont=[
    "item_price", "item_salecount", "item_score",  "store_level", "item_timestamp", "store_timestamp"
]
header_cont_user=[
    "user_age", "user_historysum"
]
header_time=[
    "timestamp"
]
header_label=[
    "is_trade"
]

# 从数据库汇总获取数据，并进行基础的归一化和独热处理
def basic_data_process():
    # 连接数据库，获取record、item、store、user四个表的信息
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
    df_cont = df_result[header_cont]
    df_cont_user = df_result[header_cont_user]
    df_id = df_result[header_id]
    df_cate = df_result[header_cate]
    df_time = df_result[header_time]
    df_label = df_result[header_label]

    # 对cont类型进行归一化处理
    scaler = preprocessing.MinMaxScaler(feature_range=(0, 1))
    scaled = scaler.fit_transform(df_cont)
    df_cont = pd.DataFrame(scaled)

    # 对cate类型进行独热编码
    df_cate = pd.get_dummies(df_cate)

    frame = [df_cont, df_cont_user, df_cate]
    X = pd.concat(frame, axis=1)
    y = df_label

    # 将结果存盘
    df_result = pd.concat([df_id, X, df_time, y], axis=1)
    # df_result.to_csv("./train_data.csv",encoding='gbk')
    return X, y


# 生成LSTM的序列化数据（t-2, t-1, t预测t）
def series_to_supervised(data, n_in=2, n_out=1, dropnan=True):
    n_vars = 1 if type(data) is list else data.shape[1]
    df = pd.DataFrame(data)
    cols, names = list(), list()
    # input sequence (t-n, ... t-1)
    for i in range(n_in, 0, -1):
        cols.append(df.shift(i))
        names += [('var%d(t-%d)' % (j + 1, i)) for j in range(n_vars)]
    # forecast sequence (t, t+1, ... t+n)
    for i in range(0, n_out):
        cols.append(df.shift(-i))
        if i == 0:
            names += [('var%d(t)' % (j + 1)) for j in range(n_vars)]
        else:
            names += [('var%d(t+%d)' % (j + 1, i)) for j in range(n_vars)]
    # put it all together
    agg = pd.concat(cols, axis=1)
    agg.columns = names
    # drop rows with NaN values
    if dropnan:
        agg.dropna(inplace=True)
    return agg


# 生成mlr——lstm的数据
def get_ml_data(timestamp=2):
    X, y = basic_data_process()
    reframed = series_to_supervised(X.values, n_in=timestamp)
    reframed = pd.DataFrame(reframed.values[:, :-1])
    y = pd.DataFrame(y.values[timestamp:])
    train_X = reframed.values
    train_X = train_X.reshape(
        (train_X.shape[0], 1, train_X.shape[1]))
    print(train_X.shape, y.values.shape)
    X = pd.DataFrame(X.values[timestamp:])
    return X, train_X, y


def main():
    # X, y = basic_data_process()
    X, train_X, y = get_ml_data(timestamp=2)
    print(X)
    print(train_X)
    print(y)

if __name__ == '__main__':
    main()