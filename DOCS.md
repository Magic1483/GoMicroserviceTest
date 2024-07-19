
# kafka 
брокер сообщений 

**использует TCP**

topic - theme or CHANNEL for messages in kafka


**kafka use case map**
```
    send msgs                   recieve msgs
    |-------->  | messages   | ------->----|
    |           |                          |
| Client | -->  | reactions  | ------> | Server |
    |           |                          |
    |-------->  | gps coords | ------->----|
    send coords                recieve coords
```
 
-----------
Проблемы http
- возможная потеря данных


# docker 

-p (--publish) docker_port:host_port
