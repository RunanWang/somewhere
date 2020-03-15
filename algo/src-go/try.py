from flask import Flask, jsonify
from flask import _request_ctx_stack
from gevent.pywsgi import WSGIServer
from gevent import monkey
import time
import train_deep_wide
from keras.layers import Input, Dense, Lambda, multiply, concatenate
from keras.layers import Dropout
from keras import backend as K
from keras import regularizers
import data_process as data_process
import h5py

monkey.patch_all()
app = Flask(__name__)
#app.lock = threading.Lock()
#app.is_training = False
 
# @app.route('/')
# def index():
#     app.lock.acquire()
#     app.running_process = app.running_process + 1
#     print("start",_request_ctx_stack._local.__ident_func__(),"running", app.running_process)
#     app.lock.release()
#     # print(_request_ctx_stack._local.__ident_func__())
#     time.sleep(30)
#     app.lock.acquire()
#     app.running_process = app.running_process - 1
#     print("end", _request_ctx_stack._local.__ident_func__(),"running",app.running_process)
#     app.lock.release()
#     return '<h1>hello</h1>'

@app.route('/status', methods=['GET'])
def index2():
    # app.lock.acquire()
    # status = app.is_training
    # app.lock.release()
    if status:
        return jsonify({"status":"training"})
    else:
        return jsonify({"status":"serving"})


@app.route('/train', methods=['POST'])
def index3():
    # app.lock.acquire()
    # if app.is_training:
    #     return jsonify({"err_code":1503, "err_msg":"Another model is training."})
    # app.is_training = True
    # app.lock.release()
    train_deep_wide.build_serve("./deep_wide.h5")
    # try:
    #     # t1 = threading.Thread(target=train_deep_wide.build_wide_n_deep, args=("./deep_wide.h5",))
    #     # t1.start()
    #     # _thread.start_new_thread( train_deep_wide.build_wide_n_deep, ("./deep_wide.h5", ) )
    #     # train_deep_wide.build_wide_n_deep("./deep_wide.h5")
    # except:
    #     return jsonify({"err_code":1003, "err_msg":"Something wrong."})
    # app.lock.acquire()
    # app.is_training = False
    # app.lock.release()
    return jsonify({"err_code":0, "err_msg":"OK, training."})


if __name__ == '__main__':
    server = WSGIServer(("0.0.0.0",12345),app)
    server.serve_forever()
    #app.run(port=12345)