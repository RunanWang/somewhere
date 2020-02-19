>使用数据集需引用文章 
F. Maxwell Harper and Joseph A. Konstan. 2015. The MovieLens Datasets: History and Context. ACM Transactions on Interactive Intelligent Systems (TiiS) 5, 4: 19:1–19:19. <https://doi.org/10.1145/2827872>

## Summary

movielens数据集是一个电影评分与电影文本标记的数据集，其中小数据集来自9742部电影的由610名用户产生的包含100836个带时间戳的评分记录、3683个文本标签。被选的用户至少对20部电影做出了评价，每个用户只包含一个用户ID信息，不包含其他信息。

movielens原始文件中包含四个文件，每个文件中的内容如下表：


| 文件名称   | 文件内容                                                 |
| ---------- | -------------------------------------------------------- |
| links.csv  | movielens的movieID与IMDB等movieID的对应关系              |
| movies.csv | movie的基本信息，包括一个movieID、类别和名称             |
| rating.csv | 评分信息，包括userID、movieID、评分（五分制0.5）和时间戳 |
| tags.csv   | 用户给出的自由文本标签                                   |

## 数据信息说明

### movie类别

电影的类别被分为：

```python
cate = [
    "Action", "Adventure", "Animation", "Children's", "Comedy", "Crime",
    "Documentary", "Drama", "Fantasy", "Film-Noir", "Horror", "Musical",
    "Mystery", "Romance", "Sci-Fi", "Thriller", "War", "Western"
]
```

## 数据处理

### movie处理

- ml_data_process_movies.py: 将movies.csv中的信息按照类别重新编组成One-hot模式，存在process_movies.csv中。

### rate处理

- ml_data_process_movie_rate.py: 统计电影的总评价数、电影平均评分存在process_statics_movies.csv中。
- ml_data_process_user_rate.py: 统计用户的总评价数、用户平均评分存在process_statics_users.csv中。

### 整合处理

- ml_data_process_final.py: 把记录和movies.csv中的电影分类信息结合，形成一张用于训练的记录表。存在process_final.csv

### 排序

- ml_data_process_sort.py: 按照uesrID进行分类，按照时间戳进行排序。存在process_sorted.csv中。

### 加标签

- ml_data_process_label.py: 去掉rating并把rating转化为label，约有0.2的评分在4分以上，可以以4分为分界线，高于4分视为1，否则视为0.存在process_labeled.csv中。

### 划分数据集

- ml_data_process_divide.py: 按照uesrID进行分类，按照时间顺序取前90%作为训练集，后10%作为测试集，（9万条训练，1万条测试），结果存在process_train.csv和process_test.csv中。去掉UserID和MovieID，