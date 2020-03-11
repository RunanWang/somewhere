from keras.models import Model
from keras.layers import Input, Dense, Lambda, multiply, concatenate
from keras.layers import Dropout
from keras import backend as K
from keras import regularizers
import data_process as data_process
import h5py


# 用于mlr模型两个部分相乘时的形状
def keras_sum_layer_output_shape(input_shape):
    # a function calculate the shape(equal to 1 in the sum func)
    shape = list(input_shape)
    assert len(shape) == 2
    shape[-1] = 1
    return tuple(shape)


# 用于mlr模型两个部分相乘时计算相乘后相加部分
def keras_sum_layer(x):
    # a function to take sum of the layers
    return K.sum(x, axis=1, keepdims=True)


# 这个函数搭建了wide&deep模型
# wide_m是MLR的分区数，默认为4，论文中是12,1时是LR。
def build_wide_n_deep(path, wide_m=4, epoch=20):
    X, y = data_process.basic_data_process()
    ## 构建wide部分，是一个MLR
    # 第一层为输入层
    input_wide = Input(shape=(X.shape[1], ))
    # 第二层为LR和权重层，采用l2正则化项
    wide_divide = Dense(wide_m, activation='softmax')(input_wide)
    wide_fit = Dense(wide_m, activation='sigmoid')(input_wide)
    # 第三层是LR和权重的乘积
    wide_ele = multiply([wide_divide, wide_fit])
    wide = Lambda(keras_sum_layer,
                output_shape=keras_sum_layer_output_shape)(wide_ele)


    ## 构建deep部分，是一个DNN
    # 第一层是输入层
    input_deep = Input(shape=(X.shape[1], ))
    # 然后是深度网络，可以叠加多层，可以加一些dropout
    deep_layer1 = Dense(32,
                        activation='relu',
                        bias_regularizer=regularizers.l2(0.01))(input_deep)
    deep_layer2 = Dropout(0.2)(deep_layer1)
    deep_layer3 = Dense(16,
                        activation='relu',
                        bias_regularizer=regularizers.l2(0.01))(deep_layer2)
    deep_layer4 = Dropout(0.2)(deep_layer3)
    deep = Dense(4, activation='sigmoid',
                bias_regularizer=regularizers.l2(0.01))(deep_layer4)
    

    ## 组合deep&wide
    coned = concatenate([wide, deep])
    # 最后用一个神经元将wide和deep部分组合
    out = Dense(1, activation='sigmoid')(coned)
    model = Model(inputs=[input_wide, input_deep], outputs=out)
    # 编译模型，选择优化函数、损失函数
    model.compile(optimizer='adam',
                loss='mean_squared_error',
                metrics=['accuracy'])
    # 进行训练并保存
    model.fit([X, X],
            y,
            epochs=epoch,
            batch_size=1)
    model.save(path)
    print("训练完毕")


def main():
    model_path = './deep_wide_model.h5'
    build_wide_n_deep(model_path)


if __name__ == '__main__':
    main()