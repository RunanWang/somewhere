import csv

# 去掉rating并把rating转化为label，约有0.2的评分在4分以上，
# 可以以4分为分界线，高于4分视为1，否则视为0.存在process_labeled.csv中。

headers = [
    "UserID", "UserNum", "UserScore", "MovieID", "MovieNum", "MovieScore",
    "Action", "Adventure", "Animation", "Children's", "Comedy", "Crime",
    "Documentary", "Drama", "Fantasy", "Film-Noir", "Horror", "Musical",
    "Mystery", "Romance", "Sci-Fi", "Thriller", "War", "Western", "timestamp",
    "label"
]

old_headers = [
    "UserID", "UserNum", "UserScore", "MovieID", "MovieNum", "MovieScore",
    "Action", "Adventure", "Animation", "Children's", "Comedy", "Crime",
    "Documentary", "Drama", "Fantasy", "Film-Noir", "Horror", "Musical",
    "Mystery", "Romance", "Sci-Fi", "Thriller", "War", "Western", "timestamp"
]

with open('process_labeled.csv', 'w', newline='') as f:
    f_csv = csv.DictWriter(f, headers)
    f_csv.writeheader()

fwrite = open('process_labeled.csv', 'a', newline='')
fw_csv = csv.DictWriter(fwrite, headers)

with open('process_sorted.csv') as f:
    data = csv.DictReader(f)
    temp_user_ID = 0
    temp_list = []
    i = 0
    for row in data:
        to_write = {}
        for item in old_headers:
            to_write[item] = row[item]
        if float(row['rating']) > 4:
            to_write['label'] = 1
        else:
            to_write['label'] = 0
        fw_csv.writerow(to_write)
        i = i + 1
        if i % 5000 == 0:
            print("已完成", i / 1000, "%")