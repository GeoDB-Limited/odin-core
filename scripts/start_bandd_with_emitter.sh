# Clear and create a new fresh DB.
dropdb my_db
createdb my_db

# Delete old kafka topic
kafka-topics --zookeeper localhost:2181 --delete --topic test
kafka-topics --zookeeper localhost:2181 --delete --topic replay

# Create kafka topic
kafka-topics --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic test
kafka-topics --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic replay

# Init table from flusher
source ./flusher/venv/bin/activate
python ./flusher/main.py init bandchain test replay --db localhost:5432/my_db


# start bandchain
bandd unsafe-reset-all --home ~/relayer/data/bandchain
bandd start --with-emitter test@localhost:9092 --rpc.laddr tcp://0.0.0.0:26657 --home ~/relayer/data/bandchain --pruning=nothing
