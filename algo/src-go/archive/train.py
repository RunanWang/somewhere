#!/usr/bin/env python
# coding: utf-8

# In[54]:


import pymongo
import pandas as pd

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
# for item in stores:
#     print(item)
df_rec = pd.DataFrame(list(records))
df_item = pd.DataFrame(list(items))
df_user = pd.DataFrame(list(users))
df_store = pd.DataFrame(list(stores))
# print(df_rec)
# print(df_item)
# print(df_user)
# print(df_store)


# In[55]:


# print(df_rec)
df_item = df_item.rename(columns={'_id':'item_id'})
df_item.set_index(['item_id'],inplace=True)
# print(df_item)
df_user = df_user.rename(columns={'_id':'user_id'})
df_user.set_index(['user_id'],inplace=True)
df_store = df_store.rename(columns={'_id':'store_id','timestamp':'store_timestamp'})
df_store.set_index(['store_id'],inplace=True)


# In[56]:


df_result = df_rec.set_index(['item_id'])
df_result = pd.concat([df_result, df_item], axis = 1, join='inner')
df_result = df_result.reset_index()
# print(df_result)
df_result.set_index(['user_id'],inplace=True)
# print(df_result)
df_result = pd.concat([df_result, df_user], axis = 1, join='inner')
df_result = df_result.reset_index()
# print(df_user)
# print(df_result)
df_result.set_index(['store_id'],inplace=True)
df_result = pd.concat([df_result, df_store], axis = 1, join='inner')
df_result = df_result.reset_index()
# df_result.drop(['item_id'])
df_result.to_csv("./result.csv",encoding='gbk')


# In[57]:


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
    "timestamp", "item_timestamp", "user_timestamp", "store_timestamp"
]
header_label=[
    "is_trade"
]
df_cont = df_result[header_cont]
df_cont_user = df_result[header_cont_user]
df_id = df_result[header_id]
df_cate = df_result[header_cate]
df_time = df_result[header_time]
df_label = df_result[header_label]


# In[58]:


from sklearn import preprocessing
scaler = preprocessing.MinMaxScaler(feature_range=(0, 1))
scaled = scaler.fit_transform(df_cont)
df_cont = pd.DataFrame(scaled)
print(df_cont)


# In[59]:


df_cate = pd.get_dummies(df_cate)
print(df_cate)
frame = [df_cont,df_cont_user, df_cate]
X = pd.concat(frame, axis=1)
y = df_label


# In[62]:


from keras.models import Model
from keras.layers import Input, Dense, Lambda, multiply
from keras import backend as K
from keras import regularizers
import utils as utils
import h5py

MODEL_PATH = './mlr_model.h5'

def keras_sum_layer_output_shape(input_shape):
    # a function calculate the shape(equal to 1 in the sum func)
    shape = list(input_shape)
    assert len(shape) == 2
    shape[-1] = 1
    return tuple(shape)


def keras_sum_layer(x):
    # a function to take sum of the layers
    return K.sum(x, axis=1, keepdims=True)

wide_m = 12
input_wide = Input(shape=(X.shape[1], ))
# 第二层为LR和权重层，采用l2正则化项
wide_divide = Dense(wide_m,
                    activation='softmax',
                    bias_regularizer=regularizers.l2(0.01))(input_wide)
wide_fit = Dense(wide_m,
                 activation='sigmoid',
                 bias_regularizer=regularizers.l2(0.01))(input_wide)
wide_ele = multiply([wide_divide, wide_fit])
out = Lambda(keras_sum_layer,
             output_shape=keras_sum_layer_output_shape)(wide_ele)
model = Model(inputs=input_wide, outputs=out)
model.compile(optimizer='adam',
              loss='mean_squared_error',
              metrics=['accuracy'])
model.fit(X,
          y,
          epochs=10,
          batch_size=2,
          callbacks=[
              utils.roc_callback(training_data=[X, y], validation_data=[X, y])
          ])
model.save(MODEL_PATH)
# model_json = model.to_json()
# with open('model.json', 'w') as file:
#     file.write(model_json)
# model.save_weights('model.json.h5')
print("训练完毕")
