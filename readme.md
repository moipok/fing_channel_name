
# Channel Name Matcher

Этот проект представляет собой утилиту для поиска строк в файле, которые содержат заданное имя в качестве подпоследовательности. Программа читает файл построчно, ищет строки, которые содержат имя в качестве подпоследовательности, и выводит их в формате, где символы имени выделены заглавными буквами.

## Как использовать

1. Убедитесь, что у вас установлен Go (версия 1.16 или выше).
2. Клонируйте репозиторий:

   ```bash
   git clone https://github.com/moipok/fing_channel_name
   ```

3. Перейдите в директорию проекта:

   ```bash
   cd fing_channel_name
   ```

4. Запустите программу:

   ```bash
   go run main.go
   ```

5. Введите имя, которое вы хотите найти в файле, и нажмите Enter.

## Пример

Если вы введете имя `петр`, программа выведет:

```
Введите имя: петр
1        автодисПЕТчеР
2        автоинсПЕкТоР
3        адаПтомЕТР
4        амПЕрмеТР
5        антроПомЕТР
6        антроПомЕТРия
7        антроПоцЕнТРизм
8        аПоцЕнТР
9        аПпрЕТиРование
10       аПпрЕТуРа
11       вПЕрёдсмоТРящий
12       генерал-инсПЕкТоР
13       гиПЕрТРофирование
14       гиПЕрТеРмия
15       гиПЕрТРофия
16       гиПоцЕнТР
```

## Структура проекта

- `main.go` — основной файл программы.
- `russian_nouns.txt` — файл со всеми русскими существительными для поиска.

