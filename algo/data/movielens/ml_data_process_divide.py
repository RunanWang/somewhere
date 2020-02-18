import csv

# 按照uesrID进行分类，按照时间顺序取前90%作为训练集，后10%作为测试集。
#（9万条训练，1万条测试）
# 结果存在process_train.csv和process_test.csv中。

def write_list(fp, write_list):
    headers = [
        "UserNum", "UserScore", "MovieNum", "MovieScore", "Action", "Adventure",
        "Animation", "Children's", "Comedy", "Crime", "Documentary", "Drama",
        "Fantasy", "Film-Noir", "Horror", "Musical", "Mystery", "Romance",
        "Sci-Fi", "Thriller", "War", "Western", "timestamp", "label"
    ]
    to_write={}
    for temp in write_list:
        for iitem in headers:
            to_write[iitem]=temp[iitem]
        fp.writerow(to_write)
    
    

old_headers = [
    "UserID", "UserNum", "UserScore", "MovieID", "MovieNum", "MovieScore",
    "Action", "Adventure", "Animation", "Children's", "Comedy", "Crime",
    "Documentary", "Drama", "Fantasy", "Film-Noir", "Horror", "Musical",
    "Mystery", "Romance", "Sci-Fi", "Thriller", "War", "Western", "timestamp",
    "label"
]

headers = [
    "UserNum", "UserScore", "MovieNum", "MovieScore", "Action", "Adventure",
    "Animation", "Children's", "Comedy", "Crime", "Documentary", "Drama",
    "Fantasy", "Film-Noir", "Horror", "Musical", "Mystery", "Romance",
    "Sci-Fi", "Thriller", "War", "Western", "timestamp", "label"
]

with open('process_train.csv', 'w', newline='') as f:
    f_csv = csv.DictWriter(f, headers)
    f_csv.writeheader()

with open('process_test.csv', 'w', newline='') as f:
    f_csv = csv.DictWriter(f, headers)
    f_csv.writeheader()

fwrite_train = open('process_train.csv', 'a', newline='')
fw_train = csv.DictWriter(fwrite_train, headers)
fwrite_test = open('process_test.csv', 'a', newline='')
fw_test = csv.DictWriter(fwrite_test, headers)

with open('process_labeled.csv') as f:
    data = csv.DictReader(f)
    temp_userID = 0
    temp_user_num = 0
    temp_train = []
    temp_test = []
    temp_count = 0
    i = 0
    for row in data:
        if temp_userID != row['UserID']:
            if temp_userID != 0:
                write_list(fw_train,temp_train)
                write_list(fw_test,temp_test)   
            temp_user_num = int(row['UserNum'])
            temp_userID = row['UserID']
            temp_train = []
            temp_test = []
            temp_count = 1
            temp_train.append(row)
        else:
            if temp_count < temp_user_num * 0.6:
                temp_train.append(row)
            else:
                temp_test.append(row)
            temp_count = temp_count + 1
        i = i + 1
        if i % 5000 == 0:
            print("已完成", i / 1000, "%")
    write_list(fw_train,temp_train)
    write_list(fw_test,temp_test) 

fwrite_train.close()
fwrite_test.close()