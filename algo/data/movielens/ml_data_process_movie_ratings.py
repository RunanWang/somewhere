import csv

# 统计四个数据作为连续的特征：电影的总评价数、电影平均评分、用户总评价数、用户平均评分
# 存在statics.csv中

headers = ["MovieID", "totalnum", "meanrate"]

movie_num = {}
movie_rate = {}

with open('ratings.csv', encoding='UTF-8') as f:
    f_csv = csv.DictReader(f)
    for row in f_csv:
        if row['movieId'] in movie_num:
            movie_num[row['movieId']] = movie_num[row['movieId']] + 1
            movie_rate[row['movieId']] = movie_rate[row['movieId']] + float(row['rating'])
        else:
            movie_num[row['movieId']] = 1
            movie_rate[row['movieId']] = float(row['rating'])

with open('process_statics_movies.csv', 'w', newline='') as f:
    f_csv = csv.DictWriter(f, headers)
    f_csv.writeheader()
    for (k, v) in movie_num.items():
        towrite = {}
        towrite['MovieID'] = k
        towrite['totalnum'] = v
        towrite['meanrate'] = movie_rate[k]/v
        f_csv.writerow(towrite)
