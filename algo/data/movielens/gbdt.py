import pandas as pd

cate_headers = [
    "Action", "Adventure",
    "Animation", "Children's", "Comedy", "Crime", "Documentary", "Drama",
    "Fantasy", "Film-Noir", "Horror", "Musical", "Mystery", "Romance",
    "Sci-Fi", "Thriller", "War", "Western"
]

cont_headers = [
    'UserNum','UserScore','MovieNum','MovieScore'
]

useless_headers =[
    "UserID", "MovieID", "timestamp"
]

def read_data(path, csv_header=0):
    data = pd.read_csv(path, header= csv_header)
    cont = data[['UserNum','UserScore','MovieNum','MovieScore']]
    cate = data[cate_headers]
    label = data['label']
    useless = data[useless_headers]
    # print(cont.head())
    # print(cate.head())
    return cont, cate, label, useless


def divide_data(data):
    cont = data[cont_headers]
    cate = data[cate_headers]
    label = data['label']
    return cont, cate, label


def cate_cal(cate, label):
    cate1_label1 = {}
    cate0_label1 = {}
    cate1_label0 = {}
    cate0_label0 = {}
    sum_score_list = {}
    for item in cate_headers:
        cate1_label1[item] = 0
        cate0_label1[item] = 0
        cate1_label0[item] = 0
        cate0_label0[item] = 0

    for index, row in cate.iterrows():
        for item in cate_headers:
            if label[index]==1:
                if row[item]==1:
                    cate1_label1[item] = cate1_label1[item] + 1
                else:
                    cate0_label1[item] = cate0_label1[item] + 1
            else:
                if row[item]==1:
                    cate1_label0[item] = cate1_label0[item] + 1
                else:
                    cate0_label0[item] = cate0_label0[item] + 1
    # print(cate1_label1)
    # print(cate0_label1)
    # print("总条数",cate0_label0['Action']+cate0_label1['Action']+cate1_label0['Action']+cate1_label1['Action'])
    # print("Action中label为1",cate1_label0['Action']+cate1_label1['Action'])
    for item in cate_headers:
        if cate0_label0[item]+cate0_label1[item] == 0:
            c0 = -1
        else:
            c0 = float(cate0_label1[item])/(cate0_label0[item]+cate0_label1[item])
        if cate1_label0[item]+cate1_label1[item] == 0:
            c1 = -1
        else:
            c1 = float(cate1_label1[item])/(cate1_label0[item]+cate1_label1[item])
        sum_score = (1-c0)*(1-c0)*cate0_label1[item] + (0-c0)*(0-c0)*cate0_label0[item] + (1-c1)*(1-c1)*cate1_label1[item] + (0-c1)*(0-c1)*cate1_label1[item]
        sum_score_list[item] = sum_score
    # print(sum_score_list)
    return sum_score_list


def cont_cal(cont, label, sum_score_list):
    item_label0_map = {}
    item_label1_map = {}
    for item in cont_headers:
        item_value_label0_map = {}
        item_value_label1_map = {}
        item_label0_map[item] = item_value_label0_map
        item_label1_map[item] = item_value_label1_map
    for index, row in cont.iterrows():    
        for item in cont_headers:
            if row[item] not in item_label0_map[item]:
                item_label0_map[item][row[item]] = 0
                item_label1_map[item][row[item]] = 0
            if label[index] == 0:
                item_label0_map[item][row[item]] = item_label0_map[item][row[item]] + 1
            else:
                item_label1_map[item][row[item]] = item_label1_map[item][row[item]] + 1
    # print(item_label0_map['UserNum'][232]+item_label1_map['UserNum'][232])
    # print(item_label1_map['UserNum'][232])
    least_ans_list={}
    for item in cont_headers:
        temp_least_score = -1
        temp_value = 0
        for keys in item_label0_map[item]:
            numleft_label1=0
            numright_label1=0
            numleft_label0=0
            numright_label0=0
            for keys2 in item_label0_map[item]:
                if keys<keys2:
                    numleft_label1 = numleft_label1+item_label1_map[item][keys2]
                    numleft_label0 = numleft_label0+item_label0_map[item][keys2]
                else:
                    numright_label1 = numright_label1+item_label1_map[item][keys2]
                    numright_label0 = numright_label0+item_label0_map[item][keys2]
            if numleft_label0+numleft_label1 == 0:
                cleft = -1
            else:
                cleft = float(numleft_label1)/(numleft_label0+numleft_label1)
            if numright_label0+numright_label1 == 0:
                cright = -1
            else:
                cright = float(numright_label1)/(numright_label0+numright_label1)
            sum_score = (1-cleft)*(1-cleft)*numleft_label1 + (0-cleft)*(0-cleft)*numleft_label0 +(1-cright)*(1-cright)*numright_label1 + (0-cright)*(0-cright)*numright_label0 
            if temp_least_score==-1 or sum_score<temp_least_score:
                temp_least_score = sum_score
                temp_value = keys
        least_ans_list[item]=temp_value
        sum_score_list[item]=temp_least_score
    # print(least_ans_list)
    # print(sum_score_list)
    return least_ans_list, sum_score_list
    

def find_divide_point(sum_score_list, least_ans_list):
    count = 0
    minitem = ''
    minscore = -1
    for item in sum_score_list:
        if sum_score_list[item]<minscore or minscore==-1:
            minitem = item
            minscore = sum_score_list[item]
    # print(minitem)
    if minitem in cont_headers:
        divide_score = least_ans_list[minitem]
    else:
        divide_score = 1
    # print(divide_score)
    sum_score_list.pop(minitem)
    # print(sum_score_list)
    return minitem, divide_score, sum_score_list



cont, cate, label, useless = read_data('./process_labeled.csv')
sum_score_list = cate_cal(cate, label)
least_ans_list, sum_score_list = cont_cal(cont, label, sum_score_list)
num = 0
df = pd.concat([cont,cate,label], axis=1)

while num < 6:
    min_item, divide_score, sum_score_list = find_divide_point(sum_score_list, least_ans_list)
    print(num, "mid-区分特征：", min_item, "区分点：", divide_score)


    df_left = df[df[min_item]<divide_score] 
    df_right = df[df[min_item]>=divide_score] 

    cont, cate, label = divide_data(df_left)
    sum_score_list_left = cate_cal(cate, label)
    least_ans_list_left, sum_score_list_left = cont_cal(cont, label, sum_score_list_left)
    min_item_left, divide_score_left, sum_score_list_left = find_divide_point(sum_score_list_left, least_ans_list_left)
    if min_item_left == min_item:
        min_item_left, divide_score_left, sum_score_list_left = find_divide_point(sum_score_list_left, least_ans_list_left)
    print(num, "left-区分特征：", min_item_left, "区分点：", divide_score_left)
    df_left_left = df_left[df_left[min_item_left]<divide_score_left] 
    df_left_right = df_left[df_left[min_item_left]>=divide_score_left] 

    cont, cate, label = divide_data(df_right)
    sum_score_list_right = cate_cal(cate, label)
    least_ans_list_right, sum_score_list_right = cont_cal(cont, label, sum_score_list_right)
    min_item_right, divide_score_right, sum_score_list_right = find_divide_point(sum_score_list_right, least_ans_list_right)
    if min_item_right == min_item:
        min_item_right, divide_score_right, sum_score_list_right = find_divide_point(sum_score_list_right, least_ans_list_right)
    print(num, "right-区分特征：", min_item_right, "区分点：", divide_score_right)
    df_right_left = df_right[df_right[min_item_right]<divide_score_right] 
    df_right_right = df_right[df_right[min_item_right]>=divide_score_right] 


    col_name0 = 'gbdt_' + str(num) +'0'
    col_name1 = 'gbdt_' + str(num) +'1'
    df_left_left.insert(len(df_left_left.columns), col_name0, 0)
    df_left_left.insert(len(df_left_left.columns), col_name1, 0)
    df_left_right.insert(len(df_left_right.columns), col_name0, 0)
    df_left_right.insert(len(df_left_right.columns), col_name1, 1)
    df_right_left.insert(len(df_right_left.columns), col_name0, 1)
    df_right_left.insert(len(df_right_left.columns), col_name1, 0)
    df_right_right.insert(len(df_right_right.columns), col_name0, 1)
    df_right_right.insert(len(df_right_right.columns), col_name1, 1)

    df = pd.concat([df_left_left, df_left_right, df_right_left, df_right_right])
    num = num + 1
    
df.sort_index(inplace=True)
df = pd.concat([df, useless], axis=1)
df.to_csv('gbdt.csv', index=0)