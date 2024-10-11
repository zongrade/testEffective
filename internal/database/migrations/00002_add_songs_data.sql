--TODO: добавить песни
-- +goose Up
-- +goose StatementBegin
INSERT INTO
    songs (group_name, song_name, link, release_date)
VALUES
    (
        'Inuma',
        'Not Good',
        'yotube.com/placeholder',
        '2000-06-27'
    );

INSERT INTO
    couplets (song_id, couplet_number, couplet_text)
VALUES
    (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 1, 'They said I have to be good
But I saw them not being good themselves
Now why should I be
How can they expect the world
To be good not being good themselves?'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 2, 'Every effort that I take to be better than I am
Makes me feel that I am fake
Is it all for my own sake?
I don''t wanna be good now
If the world''s pushing me down
I''ll resist I will break out
No doubt'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 3, 'I''m not gonna blow out my brains
Isn''t it good news?
I''ll do my best ''til I can
Death is the only truth
I''m not gonna kill all the bad
I am not good too
This world is a little bit mad
Can you just see me through?
See me through
Can you see me?'
    );

INSERT INTO
    songs (group_name, song_name, link, release_date)
VALUES
    (
        'Inuma',
        'R.37',
        'yotube.com/placeholder',
        '2000-06-27'
    );

INSERT INTO
    couplets (song_id, couplet_number, couplet_text)
VALUES
    (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 1, 'Мне стоило сбежать
Но не вернуть назад
Границы, поезда, сорвать тормоза
А тебе есть о чем рассказать? Смотри в глаза
Мне стоило сбежать от порядка, теперь играю в прятки
Я не верю тому, кто не любит рассвет
Умираю во сне, только сны без ответа
И засохшая на губах эта красная на вкус
Снова я бегу туда, где пусто'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 2, '(Пустота)'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 3, 'Нас ещё можно спасти
Если сами захотим
Куда идёшь? Некуда идти
Смотри, как горит Рим'
    );

INSERT INTO
    songs (group_name, song_name, link, release_date)
VALUES
    (
        'Inuma',
        'Выбирай',
        'yotube.com/placeholder',
        '2000-06-27'
    );

INSERT INTO
    couplets (song_id, couplet_number, couplet_text)
VALUES
    (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 1, 'Опять оставляешь пустые страницы
        Остановиться, не спать
        Тебе это только снится
        Чужие забытые лица
        Кто ты?
        Сколько лет пустоты
        Ты не знаешь куда идти'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 2, 'Выбирай свой рай
        Давно пора
        Игра не твоя, но играй'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 3, 'Даже имя твоё мне казалось как минимум странным
Ты словно солнце
Я чувствую как проникаешь в открытые раны
Дай потерять тебя
Мы прокляты что-ли
Ты хочешь исчезнуть
Но мир тонет
Я не боюсь пустоты и боли'
    );

INSERT INTO
    songs (group_name, song_name, link, release_date)
VALUES
    (
        'Б.А.У',
        'Дядяшка Сяо',
        'yotube.com/placeholder',
        '2000-06-27'
    );

INSERT INTO
    couplets (song_id, couplet_number, couplet_text)
VALUES
    (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 1, 'Жизнь в глубинке совсем не сахар
Встаю с петухами доить корову
Из развлечений – самогон и Малахов
А в местном магазине только спички и порох
Но я открыл покупки в интернете
Из далёкой страны спящего дракона
Идут посылки по всей планете
И в столицы, и в регионы'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 2, 'Пришли мне, дядюшка Сяо
Censor из Алиэкспресса
Чтоб она помогала в хозяйстве
И избавила бы от стресса
Пришли мне, дядюшка Сяо
Censor из Алиэкспресса
Чтоб она работала по блютусу
И имела бы вкус майонеза'
    );

INSERT INTO
    songs (group_name, song_name, link, release_date)
VALUES
    (
        'Б.А.У',
        'Танцы на костях',
        'yotube.com/placeholder',
        '2000-06-27'
    );

INSERT INTO
    couplets (song_id, couplet_number, couplet_text)
VALUES
    (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 1, 'Оу)
Что за страшный вой?
Что за жуткий звук?
Померк солнца свет
Затряслось все вокруг
Это мой сосед
По лестничной площадке
Подрубил сабвуфер
В своей десятке'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 2, 'Подтянулись со дворов соседних алкаши
За мерзавчиком в КБшечку ты поспеши
Сегодня снова повод выпить - праздник у Санька
Год прошел с тех пор, как от него ушла жена'
    );

INSERT INTO
    songs (group_name, song_name, link, release_date)
VALUES
    (
        'Б.А.У',
        'Потанцуй со мной',
        'yotube.com/placeholder',
        '2000-06-27'
    );

INSERT INTO
    couplets (song_id, couplet_number, couplet_text)
VALUES
    (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 1, 'Я отдыхал с другом в Таиланде
Зашли мы в барчик посидеть
Заняли место на веранде
Напитки пить, закат смотреть
Вокруг огни, играет диско
"Вон ледибой" – мне друг шепнул
Наверно местная артистка
Подумал я и подмигнул'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 2, 'Ледибой, ледибой
Потанцуй со мной!
Ледибой, ледибой
Потанцуй со мной!
Я уже чуть бухой
Потанцуй со мной!
Ледибой, ледибой
Потанцуй со мной!'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 3, 'Я подошёл, заговорил
И чую – холод по спине
Он грустный взгляд свой опустил
В тот вечер он поведал мне
Что эта жизнь ему не мила
И как он хочет на завод
Углём чтобы топить горнила
Носить спецовку круглый год'
    );

INSERT INTO
    songs (group_name, song_name, link, release_date)
VALUES
    (
        'Йорш',
        'Андерграуд',
        'yotube.com/placeholder',
        '2000-06-27'
    );

INSERT INTO
    couplets (song_id, couplet_number, couplet_text)
VALUES
    (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 1, 'Ты помнишь как весело бились витрины
Мальчишеским смехом областных магазинов
Четыре аккорда, бутылка портвейна
И кто что умеет от Горшка до Кобейна
Варёные джинсы, рваные кеды
За каждым углом гопота и скинхеды
И первая ночь в 16 лет в отделении
Зато мы узнали кто мы на самом деле'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 2, 'Но когда я умру, проиграю свой раунд
Меня закопайте в сырой андеграунд
Не в ад и не в рай, там закрыты все двери
Давай проверим, давай все же проверим
Когда я умру, проиграю свой раунд
Меня закопайте в сырой андеграунд
Не в ад и не в рай, там закрыты все двери
Кинь горсть земли в последний мой оупен-эйр'
    );

INSERT INTO
    songs (group_name, song_name, link, release_date)
VALUES
    (
        'Йорш',
        'Половинки',
        'yotube.com/placeholder',
        '2000-06-27'
    );

INSERT INTO
    couplets (song_id, couplet_number, couplet_text)
VALUES
    (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 1, 'На серых домах цветные картинки
Ну что, по целой или по половинке
Под этим окном расцветали пионы
А сейчас только маски, быки и загоны'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 2, 'На серых домах цветные картинки
Ну что, по целой или по половинке
Под этим окном мы с тобой целовались
Странно — маски исчезли, а загоны остались'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 3, 'Давай разрушим боль, давай притупим психоз
Чёрно-серый пейзаж из КПП и берёз
Снова за гаражи поползла детвора
Обойдём этот дом, смотри — стоят мусора
Возьми за руку, надоел такой рай
Здесь только грязь, здесь вечный февраль
Противный мелкий дождь на узорах лица
И люди в штатском пасут у крыльца'
    );

INSERT INTO
    songs (group_name, song_name, link, release_date)
VALUES
    (
        'Йорш',
        'Нарисуй мне счастье',
        'yotube.com/placeholder',
        '2000-06-27'
    );

INSERT INTO
    couplets (song_id, couplet_number, couplet_text)
VALUES
    (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 1, 'У нас же всё так!
Дороги ремонтируют, когда президент приезжает
Мост строят, когда саммит в городе
На курорте порядок навести — только если Олимпиада к нам приедет
Только вот в этом случае
Ну можно же что-то сделать, вот?
Ну не... да не страна, ну декорация какая-то!'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 2, 'Нарисуй мне счастье, нарисуй свободу!
Нарисуй мне пищу, нарисуй мне воду!
Нарисуй больницы, школы и дороги!
Над страной брезента формалиновые боги!'
    );

INSERT INTO
    songs (group_name, song_name, link, release_date)
VALUES
    (
        'KDRR',
        'Ева Эльфи',
        'yotube.com/placeholder',
        '2000-06-27'
    );

INSERT INTO
    couplets (song_id, couplet_number, couplet_text)
VALUES
    (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 1, 'Он мог бы кем угодно быть
Но стал в Сургуте швы варить
Из институтов - армия
А в ксиве штамп - статья'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 2, 'Вагон, шконарь, в окно фонарь
Вокруг калейдоскоп из харь
Сто граммов, чтоб не околеть
И перед сном тайком смотреть'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 3, 'На Еву Элфи
На Еву Элфи
На Еву Элфи
Я осуждаю, но смотрю
На Еву Элфи
На Еву Элфи
На Еву Элфи'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 4, 'А в Омске ждёт его она
Та, что судьбой поручена
Она тут самая ващще
Амбассадор борщей'
    );

INSERT INTO
    songs (group_name, song_name, link, release_date)
VALUES
    (
        'KDRR',
        'Солнышко',
        'yotube.com/placeholder',
        '2000-06-27'
    );

INSERT INTO
    couplets (song_id, couplet_number, couplet_text)
VALUES
    (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 1, 'Маленькие, чёрные мысли в моей голове
Всё не оно, вкривь да вкось, да как-то не так
Во дворе заржавели качели и не по себе
Зреет, звереет вокруг животный бардак'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 2, 'Фонари разобьют нашу улицу вдребезги, чтобы мы не вернулись домой
Последняя песня воскресного вечера
Последний приют, где остыло твоё'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 3, 'Солнышко в руках
И венец из звёзд в небесах
И с других планет все видят нас
Нам так глубоко с тобой плевать на это
Где-то над землёй мы парим с тобой в облаках
Где-то в перевернутых снах
Мы торчим на счастье и теплом ветре'
    );

INSERT INTO
    songs (group_name, song_name, link, release_date)
VALUES
    (
        'Nagart',
        'Пугало',
        'yotube.com/placeholder',
        '2000-06-27'
    );

INSERT INTO
    couplets (song_id, couplet_number, couplet_text)
VALUES
    (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 1, 'Жутко всё было так
Есть на кладбище овраг
Столько молвит о нём народ
Только Леший разберёт'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 2, 'Жутко всё было так
Есть на кладбище овраг
Столько молвит о нём народ
Только Леший разберёт'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 3, 'Кстати, он свидетелем был
Как бродяга старуху бранил
Пугалом называл
Громко, дико хохотал'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 4, '"Погляди на меня!", — кричал
"Я самих королей встречал
На балах у князей кутил
В Париже и Лондоне жил!"'
    );

INSERT INTO
    songs (group_name, song_name, link, release_date)
VALUES
    (
        'Nagart',
        'Умрун',
        'yotube.com/placeholder',
        '2000-06-27'
    );

INSERT INTO
    couplets (song_id, couplet_number, couplet_text)
VALUES
    (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 1, 'В грязном сумраке таверны
В тусклом свете трёх свечей
Скоморох довольно скверный
И в заплатанном плаще'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 2, 'Лихо пальцами терзал
Остатки ржавых струн
Хриплым голосом вещал
Что копит силы чародей Умрун'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 3, 'За межой гнилого леса
Людям смертным хода нет
Будь ты конный или пеший
Растворится человек'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 4, 'В чародействе паутины
Вдоль гнилых тенистых троп
В мороке бесовской силы
Тлеет под ковром лесных цветов'
    );

INSERT INTO
    songs (group_name, song_name, link, release_date)
VALUES
    (
        'Nagart',
        'Моя тюрьма',
        'yotube.com/placeholder',
        '2000-06-27'
    );

INSERT INTO
    couplets (song_id, couplet_number, couplet_text)
VALUES
    (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 1, 'Беснуется народ по жизни от того
Что мало в ней хлопот, накрыла лень его!
Король придумал: "Надо взять идею, как во сне!
Законы нужно сочинять, в каменной тюрьме!"'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 2, 'Страданья и кошмар
Палач раздует жар
Когда услышит боль
Работает король!'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 3, 'Все узники не спят: введу налог на сон
Работать должен раб, он для того рождён!
За дверью кто-то закричал – несчастный там лишился глаз
Король, подумав, подписал о смертной казни указ!'
    );

INSERT INTO
    songs (group_name, song_name, link, release_date)
VALUES
    (
        'Norma Tale',
        'Кай и герда',
        'yotube.com/placeholder',
        '2000-06-27'
    );

INSERT INTO
    couplets (song_id, couplet_number, couplet_text)
VALUES
    (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 1, 'Говорят, смотря в глаза, видишь зеркало души
Но в её глазах осколки, скрытые за стены ширм
Она прячет своё зеркало за прожитую жизнь
Вместо блеска от стекла, я вижу только миражи'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 2, 'В нашей сказке злобный тролль закинул каждому из нас
По куску кривого зеркала в какой-нибудь из глаз
Наши розы расцветут, но видно только не сейчас
А пока друг другу в окна, смотрим мы в последний раз'
    );

INSERT INTO
    songs (group_name, song_name, link, release_date)
VALUES
    (
        'Lascala',
        'Agonia',
        'yotube.com/placeholder',
        '2000-06-27'
    );

INSERT INTO
    couplets (song_id, couplet_number, couplet_text)
VALUES
    (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 1, 'На холодном балконе в доме
На грани шторма
У меня паранойя
В двое, в трое выше нормы'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 2, 'Я посылаю сигналы тревоги
В каждом начале своих монологов
Если ты мне не дашь свою ладонь
Велю открыть огонь'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 3, 'Кричу "Madonna Mia"
Огонь моя стихия
Волной смертельного тепла
Я все сожгу до тла'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 4, 'На дешевой картине
В ныне
Чужой гостиной
Я стою на вершине
Мира
Мне невыносимо'
    );

INSERT INTO
    songs (group_name, song_name, link, release_date)
VALUES
    (
        'Lascala',
        'Реванш',
        'yotube.com/placeholder',
        '2000-06-27'
    );

INSERT INTO
    couplets (song_id, couplet_number, couplet_text)
VALUES
    (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 1, 'И будет снег
Накроет белым полотном ночной побег
Удача - скользкий мячик выпрыгнет из рук
И даже, самый в мире близкий человек исчезнет вдруг'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 2, 'И сыпались на части все остатки счастья
Без жалости и сострадания
Подкидывала кости
Разжигала злостью
Минувшие воспоминания
Ха-ха-у
Так долго горевала, что в своих слезах тону
И как невыносимо покидают силы'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 3, 'Спорим, мой ручей столкнётся с морем
И затянет горизонт
Туманом моих страхов, брошенных за борт
А спорим, виноватых будет двое
Корабли вернутся в порт
На разные причалы
Ты начнёшь сначала, я - наоборот'
    );

INSERT INTO
    songs (group_name, song_name, link, release_date)
VALUES
    (
        'Lascala',
        'Против ветра',
        'yotube.com/placeholder',
        '2000-06-27'
    );

INSERT INTO
    couplets (song_id, couplet_number, couplet_text)
VALUES
    (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 1, 'На полной скорости
Если на спидометре ноль
Давай, вставай
Твой динамик - мой алкоголь
Скорей, включай
Эта история
По траектории
Поднимает облаком пыль
И пока не кончится боль
Ещё сто миль
Дети лета против ветра'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 2, 'На полной скорости
У края пропасти
Будет страшно как ни крути
Но без сомнения
Ждут приключения
Руки вверх и к небу лети
Всё выше и выше
Худшее уже позади
На полной скорости
У края пропасти
Разойдутся наши пути
И мне придется уйти'
    );

INSERT INTO
    songs (group_name, song_name, link, release_date)
VALUES
    (
        'Корsика',
        'Белое небо',
        'yotube.com/placeholder',
        '2000-06-27'
    );

INSERT INTO
    couplets (song_id, couplet_number, couplet_text)
VALUES
    (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 1, 'Ты идёшь босиком по краю
Обожжённого градом поля
И как будто необитаем
Пред тобою летит рассвет'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 2, 'Ты их видишь, и точно знаешь
На тебя смотрят пулемёты
Но идёшь — и рассвет рисует
Твой танцующий силуэт'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 3, 'Не оглянешься — за тобою
Чьё-то небо и чья-то правда
Чьи-то руки, сердца и счастье
Чья-то воля и чья-то жизнь'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 4, 'Ты идёшь по траве прожжённой
Ты идёшь и стоит колонна
И кто крикнуть "Огонь" им должен
Почему-то ещё молчит'
    );

INSERT INTO
    songs (group_name, song_name, link, release_date)
VALUES
    (
        'Корsика',
        'Последний шанс',
        'yotube.com/placeholder',
        '2000-06-27'
    );

INSERT INTO
    couplets (song_id, couplet_number, couplet_text)
VALUES
    (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 1, 'Карточный мир
Вот бы и всё на воздух
К чёрту, я так серьёзно, я не шутил'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 2, 'Блики стены, мёртвых сердец осколки
Только не вслух, не в голос
Только не мы'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 3, 'Граница здесь, граница там
Без разницы, не вариант
Ты ни при чём, ты ни о чём
Все ждут, и мы ждём'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 4, 'Ну, а реальный шанс всего один
Нам выпадает только эта жизнь
И каждый день уносит за собой
Боль не свершений и не данный бой'
    );

INSERT INTO
    songs (group_name, song_name, link, release_date)
VALUES
    (
        'Корsика',
        'Сенсей',
        'yotube.com/placeholder',
        '2000-06-27'
    );

INSERT INTO
    couplets (song_id, couplet_number, couplet_text)
VALUES
    (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 1, 'Он был старше тебя на вечность,
Он в тебе открывал секреты,
На воде зажигая свечи,
Ты - алтарь, говорил тебе он.
Его шепот жег твое имя,
Не жалея и не скрывая,
На границе ночной пустыни
Он тебе выдал эту тайну.'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 2, 'Твой последний урок -
Что победа на вкус,
Будто кровь из разбитой губы.
Твой последний урок -
Если нужен твой ход,
То слова - уже вряд ли нужны.'
    ), (
        (
            SELECT
                id
            FROM
                songs
            ORDER BY
                id DESC
            LIMIT
                1
        ), 3, 'Небо плакало и смеялось,
Облака и песок мешая,
Ты не видела, как ушел он -
Но уже почему-то знала...
Не сгибаемый, не прощенный,
Не разбитый о камни солнца,
Оставляя тебе то небо,
Куда он уже не вернется...'
    );

-- +goose StatementEnd