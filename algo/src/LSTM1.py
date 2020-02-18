from keras.models import Sequential
from keras.layers import Dense
from keras.layers import LSTM, Dropout
from keras.callbacks import TensorBoard
import preprocess as preprocess
import utils as utils
from keras import regularizers

train_X, train_y = preprocess.get_LSTM_data(path="./process_train.csv")
test_X, test_y = preprocess.get_LSTM_data()
model = Sequential()
model.add(
    LSTM(24,
         input_shape=(train_X.shape[1], train_X.shape[2]),
         ))
model.add(Dropout(0.3))
model.add(Dense(1,
          activation='sigmoid'))
model.compile(loss='mean_squared_error',
              optimizer='adam',
              metrics=['accuracy', utils.rmse])
model.fit(train_X,
          train_y,
          epochs=100,
          batch_size=10,
          callbacks=[
              utils.roc_callback(training_data=[train_X, train_y],
                                 validation_data=[test_X, test_y]),
              TensorBoard(log_dir='final/{}'.format("lstm_mse_nol2_24_0.2"))
          ])
loss , accuracy, rmse = model.evaluate(test_X, test_y)
print('Accuracy: %.2f %%' % (accuracy * 100))
print('RMSE: %.2f %%' % (rmse * 100))
print('Loss: %.2f %%' % (loss * 100))