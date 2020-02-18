import pandas as pd
from sklearn import preprocessing


# 读入数据，path为文件相对路径，csv_header为是否略过第一行，默认0
def read_data(path, csv_header=0):
    data_train = pd.read_csv(path, header=csv_header)
    # 这里是前4列是连续变量，4-22是离散变量，23是timestamp，最后是label
    cont = data_train.values[:, 0:4]
    cate = data_train.values[:, 4:34]
    label = data_train['label']
    return cont, cate, label


# 对连续变量做归一化处理
def preprocess_normal(cont):
    scaler = preprocessing.MinMaxScaler(feature_range=(0, 1))
    scaled = scaler.fit_transform(cont)
    return scaled


# 整合连续变量与离散变量形成输入
def preprocess_merge(cont, cate):
    df_cont = pd.DataFrame(cont)
    df_cate = pd.DataFrame(cate)
    frame = [df_cont, df_cate]
    X = pd.concat(frame, axis=1)
    return X


# 整合连续变量与离散变量与标签形成输入
def preprocess_LSTM_merge(cont, cate, label):
    df_cont = pd.DataFrame(cont)
    df_cate = pd.DataFrame(cate)
    frame = [df_cate, label]
    X = pd.concat(frame, axis=1)
    return X


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


def get_LSTM_data(path="./process_test.csv", timestamp=2):
    cont, cate, y = read_data(path)
    cont = preprocess_normal(cont)
    X = preprocess_LSTM_merge(cont, cate, y)
    reframed = series_to_supervised(X.values, n_in=timestamp)
    reframed = pd.DataFrame(reframed.values[:, :-1])
    y = pd.DataFrame(y.values[2:])
    train_X = reframed.values
    train_X = train_X.reshape(
        (train_X.shape[0], 1, train_X.shape[1]))
    print(train_X.shape, y.values.shape)
    return train_X, y



def get_ml_data(path="./process_test.csv", timestamp=2):
    cont, cate, y = read_data(path)
    cont = preprocess_normal(cont)
    X = preprocess_LSTM_merge(cont, cate, y)
    reframed = series_to_supervised(X.values, n_in=timestamp)
    reframed = pd.DataFrame(reframed.values[:, :-1])
    y = pd.DataFrame(y.values[timestamp:])
    train_X = reframed.values
    train_X = train_X.reshape(
        (train_X.shape[0], 1, train_X.shape[1]))
    print(train_X.shape, y.values.shape)
    X = preprocess_merge(cont, cate)
    X = pd.DataFrame(X.values[timestamp:])
    return X, train_X, y


# test
def main():
    print("基础检查")
    cont, cate, y = read_data("./process_test.csv")
    cont = preprocess_normal(cont)
    X = preprocess_merge(cont, cate)
    print(X)
    print(y)
    print("基础检查完成")
    print("LSTM数据生成检查")
    X = preprocess_LSTM_merge(cont, cate, y)
    reframed = series_to_supervised(X.values)
    reframed = pd.DataFrame(reframed.values[:, :-1])
    print(reframed)
    y = pd.DataFrame(y.values[2:])
    print(y)
    print("LSTM检查完成")
    get_LSTM_data()


if __name__ == '__main__':
    main()