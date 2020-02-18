from keras.models import Model
from keras.layers import Input, Dense, Lambda, multiply
from keras import backend as K
from keras import regularizers
import utils as utils
import preprocess as preprocess
from keras.callbacks import TensorBoard


def keras_sum_layer_output_shape(input_shape):
    # a function calculate the shape(equal to 1 in the sum func)
    shape = list(input_shape)
    assert len(shape) == 2
    shape[-1] = 1
    return tuple(shape)


def keras_sum_layer(x):
    # a function to take sum of the layers
    return K.sum(x, axis=1, keepdims=True)

# 准备训练集数据
cont, cate, y = preprocess.read_data('process_train.csv')
cont = preprocess.preprocess_normal(cont)
X = preprocess.preprocess_merge(cont, cate)
# 准备测试集数据（验证数据也用测试集数据）
test_cont, test_cate, test_y = preprocess.read_data('process_test.csv')
test_cont = preprocess.preprocess_normal(test_cont)
test_X = preprocess.preprocess_merge(test_cont, test_cate)
# 设置MLR的分区数，默认为12.
wide_m = 1
# 第一层为输入层
input_wide = Input(shape=(X.shape[1], ))
# 第二层为LR和权重层，采用l2正则化项
# wide_divide = Dense(wide_m,
#                     activation='softmax',
#                     bias_regularizer=regularizers.l2(0.01))(input_wide)
wide_fit = Dense(wide_m,
                 activation='sigmoid',
                 bias_regularizer=regularizers.l2(0.01))(input_wide)
# 第三层是乘积
# wide_ele = multiply([wide_divide, wide_fit])
# out = Lambda(keras_sum_layer,
#              output_shape=keras_sum_layer_output_shape)(wide_ele)
# 编译模型
model = Model(inputs=input_wide, outputs=wide_fit)
model.compile(optimizer='adam',
              loss='binary_crossentropy',
              metrics=['accuracy', utils.rmse])
# 喂数据
model.fit(X,
          y,
          epochs=150,
          batch_size=10,
          callbacks=[
              utils.roc_callback(training_data=[X, y], validation_data=[test_X, test_y]),
              TensorBoard(log_dir='mytensorboard/{}'.format("lr"))
          ])
# 评估
_, accuracy, rmse = model.evaluate(test_X, test_y)
print('Accuracy: %.2f %%' % (accuracy * 100))
print('RMSE: %.2f %%' % (rmse * 100))

