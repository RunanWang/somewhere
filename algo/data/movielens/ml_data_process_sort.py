import csv


def sort_rows(temp_list):
    num = len(temp_list)
    i = 0
    while i < num:
        j = i + 1
        while j < num:
            if temp_list[i]['timestamp'] > temp_list[j]['timestamp']:
                temprow = temp_list[i]
                temp_list[i] = temp_list[j]
                temp_list[j] = temprow
            j = j + 1
        i = i + 1
    return temp_list


# ml_data_process_sort.py: 按照uesrID进行分类，
# 按照时间戳进行排序。存在process_sorted.csv中。

headers = [
    "UserID", "UserNum", "UserScore", "MovieID", "MovieNum", "MovieScore",
    "Action", "Adventure", "Animation", "Children's", "Comedy", "Crime",
    "Documentary", "Drama", "Fantasy", "Film-Noir", "Horror", "Musical",
    "Mystery", "Romance", "Sci-Fi", "Thriller", "War", "Western", "timestamp",
    "rating"
]

with open('process_sorted.csv', 'w', newline='') as f:
    f_csv = csv.DictWriter(f, headers)
    f_csv.writeheader()

fwrite = open('process_sorted.csv', 'a', newline='')
fw_csv = csv.DictWriter(fwrite, headers)

with open('process_final.csv') as f:
    data = csv.DictReader(f)
    temp_user_ID = 0
    temp_list = []
    i = 0
    for row in data:
        if temp_user_ID != row['UserID']:
            if temp_user_ID != 0:
                temp_list = sort_rows(temp_list)
                fw_csv.writerows(temp_list)
            temp_user_ID = row['UserID']
            temp_list = []
            temp_list.append(row)
        else:
            temp_list.append(row)
        i = i + 1
        if i % 5000 == 0:
            print("已完成", i / 1000, "%")
    temp_list = sort_rows(temp_list)
    fw_csv.writerows(temp_list)
