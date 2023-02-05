# RestAPi

## Пагинация

{
    limit 5
    offset 0
    count
}


/todo?posion=0?offset=10

{
    data: [
        {

        },
        {

        }
    ],
    pagination{
        limit: 5
        offset: 0
        count:100
    }
}


## Сортировка
/todo?sort=filde&operator=[ASC | DESC]