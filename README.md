### Kafka CLI
1. Create topic
kafka-topics --bootstrap-server=localhost:9093 --topic=topic1 --create

2. Set consumer subscribe topic
- Subscribe one topic:
kafka-console-consumer --bootstrap-server=localhost:9093 --topic=message

- Subscribe multiple topic:
kafka-console-consumer --bootstrap-server=localhost:9093 --include="message1|message2"

3. Set producer publish topic
kafka-console-producer --bootstrap-server=localhost:9093 --topic=message

4. Create consumer group 
kafka-console-consumer --bootstrap-server=localhost:9093 --topic=message --group=message1

5. List topics
kafka-topics --bootstrap-server=localhost:9093 --list

6. Describe topic
kafka-topics --bootstrap-server=localhost:9093 --describe

7. List consumer group
kafka-consumer-groups --bootstrap-server=localhost:9093 --list


