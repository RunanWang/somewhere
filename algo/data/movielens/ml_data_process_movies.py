import csv
# 将movies.csv中的信息按照类别重新编组成One-hot模式，存在process_movies.csv中。
cate = [
    "Action", "Adventure", "Animation", "Children's", "Comedy", "Crime",
    "Documentary", "Drama", "Fantasy", "Film-Noir", "Horror", "Musical",
    "Mystery", "Romance", "Sci-Fi", "Thriller", "War", "Western"
]

headers = [
    "MovieID", "Action", "Adventure", "Animation", "Children's", "Comedy",
    "Crime", "Documentary", "Drama", "Fantasy", "Film-Noir", "Horror",
    "Musical", "Mystery", "Romance", "Sci-Fi", "Thriller", "War", "Western"
]

with open('process_movies.csv', 'w', newline='') as f:
    f_csv = csv.DictWriter(f, headers)
    f_csv.writeheader()

fwrite = open('process_movies.csv', 'a', newline='')
fw_csv = csv.DictWriter(fwrite, headers)

with open('movies.csv', encoding='UTF-8') as f:
    f_csv = csv.reader(f)
    movienum = 0
    for row in f_csv:
        if row[0] == "movieId":
            continue
        rowcate = row[2].split("|")
        movienum = movienum + 1
        i = 0
        moviedetail = {}
        moviedetail['MovieID'] = row[0]
        while i < len(cate):
            if cate[i] in rowcate:
                moviedetail[cate[i]] = 1
            else:
                moviedetail[cate[i]] = 0
            i = i + 1
        fw_csv.writerow(moviedetail)
