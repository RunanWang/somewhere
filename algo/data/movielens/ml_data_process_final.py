import csv

# 把记录和movies.csv中的电影分类信息结合，形成一张用于训练的记录表
# 按照user-timestamp排序，存在process_final.csv

headers = [
    "UserID", "UserNum", "UserScore", "MovieID", "MovieNum", "MovieScore",
    "Action", "Adventure", "Animation", "Children's", "Comedy", "Crime",
    "Documentary", "Drama", "Fantasy", "Film-Noir", "Horror", "Musical",
    "Mystery", "Romance", "Sci-Fi", "Thriller", "War", "Western", "timestamp",
    "rating"
]

cate = [
    "Action", "Adventure", "Animation", "Children's", "Comedy", "Crime",
    "Documentary", "Drama", "Fantasy", "Film-Noir", "Horror", "Musical",
    "Mystery", "Romance", "Sci-Fi", "Thriller", "War", "Western"
]

with open('process_final.csv', 'w', newline='') as f:
    f_csv = csv.DictWriter(f, headers)
    f_csv.writeheader()

f1 = open('process_rating.csv', 'r')
data_rate = csv.DictReader(f1)
print("read finished!")

fwrite = open('process_final.csv', 'a', newline='')
fw_csv = csv.DictWriter(fwrite, headers)

i = 0
for row in data_rate:
    i = i + 1
    towrite = {}
    towrite['UserID'] = row['UserID']
    towrite['MovieID'] = row['MovieID']
    # if float(row['rating']) > 4:
    #     towrite['label'] = 1
    # else:
    #     towrite['label'] = 0
    towrite['rating'] = row['rating']
    towrite['timestamp'] = row['timestamp']
    f2 = open('process_movies.csv', 'r')
    data_movies = csv.DictReader(f2)
    for temprow in data_movies:
        if temprow['MovieID'] == row['MovieID']:
            for tempcate in cate:
                towrite[tempcate] = temprow[tempcate]
            break
    f2.close()
    f3 = open('process_statics_movies.csv', 'r')
    data_smovies = csv.DictReader(f3)
    for temprow in data_smovies:
        if temprow['MovieID'] == row['MovieID']:
            towrite['MovieNum'] = temprow['totalnum']
            towrite['MovieScore'] = temprow['meanrate']
            break
    f3.close()
    f4 = open('process_statics_users.csv', 'r')
    data_susers = csv.DictReader(f4)
    for temprow in data_susers:
        if temprow['UserID'] == row['UserID']:
            towrite['UserNum'] = temprow['totalnum']
            towrite['UserScore'] = temprow['meanrate']
            break
    f4.close()
    fw_csv.writerow(towrite)
    if i % 1000 == 0:
        print("已完成", i/1000, "%")
