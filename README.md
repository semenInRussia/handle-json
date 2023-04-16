# Handle-Json

15 апреля Tinkoff провёл отборочный тур своего кубка, который они
назвали красивым именем --- _It's Tinkoff cup_.  Любой желающий мог
выбрать один из нескольких "треков" (backend, frontend, SRE и другие).
Я не постеснялся и записался на трек *backend*, в котором я, жёстко
просрав все сроки, не смог сделать наилегчайшее задание, с таким
условие (примерно, ведь Tinkoff не даёт мне посмотреть сданные
задания, говорит, что срок --- всё!):

> Вам дан JSON, с ключом `data`, значение которого --- массив из
> объектов, каждый из них либо: `payment`, тогда его `type` равен
> `payment`, ещё у него есть `amount` и `id`; или `address` и его
> `type` равен `address`, ещё у него есть `address` и `id`; или (самый
> большой из этих) `user`, у него есть типичный `id`, `firstname` и
> `lastname`.  Необходимо вернуть JSON, который массив из `user`-ов, у
> которого помимо `id`, `firstname` и `lastname` есть поля `payments`
> и `addresses`, которые содержать `payment`-ы и `address`-ы
> объеденённые `user`-ом общим `id`

То есть из такого:

```json
{
  "data": [
    {
      "type": "address",
      "id": 1,
      "address": "Baker street, 222"
    },
    {
      "type": "user",
      "id": 1,
      "firstname": "Bean",
      "lastname": "Dremen"
    },
    {
      "type": "payment",
      "id": 1,
      "amount": 2
    },
    {
      "type": "payment",
      "id": 2,
      "amount": 3
    },
    {
      "type": "user",
      "id": 2,
      "firstname": "Bob",
      "lastname": "Odenkirk"
    }
  ]
}
```

Надо получить:

```json
[
  {
    "id": 1,
    "firstname": "Bean",
    "lastname": "Dremen",
    "payments": [
      {
        "id": 1,
        "amount": 2
      }
    ],
    "addresses": [
      {
        "address": "Baker street, 222",
        "id": 1
      }
    ]
  },
  {
    "id": 2,
    "firstname": "Bob",
    "lastname": "Odenkirk",
    "payments": [
      {
        "id": 2,
        "amount": 3
      }
    ],
    "addresses": []
  }
]
```

Когда время закончилось, а я это не успел, то я решил написать решения
на как можно больше языках.  Для каждого языка есть папка с именем
языка, в котором и лежить весь исходный код
