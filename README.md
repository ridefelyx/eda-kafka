# Kafka POC
# Setup
## Install Kafka
https://archive.apache.org/dist/kafka/3.0.0/kafka_2.13-3.0.0.tgz

extract and move to extracted location

### Start Zookeeper
`bin/zookeeper-server-start.sh config/zookeeper.properties`

### Create the log directories
`mkdir /tmp/kafka-logs`

### Start Kafka broker(s)
`bin/kafka-server-start.sh config/server.properties`

for multiple brokers go to config/server.properties, copy it to as many files as you want brokers, and these 3 params should be unique in each file:
- broker.id=0
- listeners=PLAINTEXT://:9092
- log.dirs=/tmp/kafka-logs

### Create topic
`bin/kafka-topics.sh --create --topic my-kafka-topic --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1` 
- partitions : how many brokers you want your data to be split between
- replication-factor : how many copies of you data you want (max is the number of brokers)
- bootstrap-server : points to the address of any one of our active Kafka brokers, it doesn’t matter which one you choose

list all topics:`bin/kafka-topics.sh --list --bootstrap-server localhost:9092`

details on topic: `bin/kafka-topics.sh --describe --topic my-kafka-topic --bootstrap-server localhost:9093`

delete topic: `bin/kafka-topics.sh --bootstrap-server localhost:9092 --delete --topic topic1`

### Start producer and send a message
`bin/kafka-console-producer.sh --broker-list localhost:9092,localhost:9093 --topic my-kafka-topic`

### Start consumer
`bin/kafka-console-consumer.sh --bootstrap-server localhost:9093 --topic my-kafka-topic --from-beginning`

# Pros and cons

## Pros
- easy to set up locally, extensive and easy to follow official documentation
- lots of tutorials and online materials for different problems that might occur
- several go libraries with good documentation and regular maintenance
- automatic topic creation can be enabled or disabled
- support avro schema

## Cons
- message ordering within partitions only
- extensive features and available materials can be overwhelming for beginners (distinguishing what is important and what isn't)

	