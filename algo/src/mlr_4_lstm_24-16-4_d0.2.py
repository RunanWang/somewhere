from keras.models import Model
from keras.layers import Input, Dense, Lambda, multiply, concatenate
from keras.layers import LSTM, Dropout
from keras import backend as K
from keras import regularizers
import pandas as pd
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
train_X_mlr, train_X_lstm, y = preprocess.get_ml_data(
    path="./process_train.csv")
# 准备测试集数据（验证数据也用测试集数据）
test_X_mlr, test_X_lstm, test_y = preprocess.get_ml_data(
    path="./process_test.csv")
# 设置MLR的分区数，默认为12.
wide_m = 4
# 第一层为输入层
input_wide = Input(shape=(train_X_mlr.shape[1], ))
# 第二层为LR和权重层，采用l2正则化项
wide_divide = Dense(wide_m, activation='softmax')(input_wide)
#bias_regularizer=regularizers.l2(0.01)
wide_fit = Dense(wide_m, activation='sigmoid')(input_wide)

# 第三层是乘积
wide_ele = multiply([wide_divide, wide_fit])
wide = Lambda(keras_sum_layer,
              output_shape=keras_sum_layer_output_shape)(wide_ele)

# 构建deep部分
input_deep = Input(shape=(train_X_lstm.shape[1], train_X_lstm.shape[2]))
deep_layer1 = LSTM(24)(input_deep)
deep_layer2 = Dropout(0.2)(deep_layer1)
deep_layer3 = Dense(16,
                    activation='relu',
                    bias_regularizer=regularizers.l2(0.01))(deep_layer2)
deep_layer4 = Dropout(0.2)(deep_layer3)
deep = Dense(4, activation='sigmoid',
             bias_regularizer=regularizers.l2(0.01))(deep_layer4)
# 组合deep&wide
coned = concatenate([wide, deep])
out = Dense(1, activation='sigmoid')(coned)
model = Model(inputs=[input_wide, input_deep], outputs=out)
model.compile(optimizer='adam',
              loss='mean_squared_error',
              metrics=['accuracy', utils.rmse])
# mean_squared_error binary_crossentropy
model.fit([train_X_mlr, train_X_lstm],
          y,
          epochs=100,
          batch_size=10,
          callbacks=[
              utils.roc_callback(training_data=[[train_X_mlr, train_X_lstm],
                                                y],
                                 validation_data=[[test_X_mlr, test_X_lstm],
                                                  test_y]),
              TensorBoard(log_dir='final/{}'.format("mlr_8_lstm_32-16-8_d0.2"))
          ])
loss, accuracy, rmse = model.evaluate([test_X_mlr, test_X_lstm], test_y)
print('Accuracy: %.2f %%' % (accuracy * 100))
print('RMSE: %.2f %%' % (rmse * 100))
print('Loss: %.2f %%' % (loss * 100))