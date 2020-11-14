import pandas as pd
import pymongo
from bson.objectid import ObjectId


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
        print(item['item_id'])
        if item['is_trade'] == 1:
            idToFreq[item['item_id']] = idToFreq[item['item_id']] + 1
            total = total + 1
    
    for k in idToFreq.keys():
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
        ans[k.__str__()] = idToScore[k]/total

    print(ans)


m = getItemCFMatrix()
# print(m)
getUserScore("5e9bf462b599f94b7a941436", m)
