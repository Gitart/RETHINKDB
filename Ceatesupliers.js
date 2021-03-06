
r.dbCreate("Headoffice");
r.db("Headoffice").tableCreate("S_SUPPLIERS") ;
r.table("S_SUPPLIERS").insert({id:1,
	                           ID_CONTRACTORS: 1,
                               NAME:         "Name Contractors",
                               NAME_EXT:     "Name Contractors EX", 
                               IS_FINAL:     1 ,
                               IS_SELLER:    1 ,  // Признак поставщика
                               IS_CUSTOMER   1 ,  // Признак покупателя 
                               ID_REITING    1 ,  // Рейтинг поставщика по 5 бальной системе
                               NUMBER_TAX    "" , //VARCHAR(255), -- Номер свидетельства
                               EMAIL:        "" , //VARCHAR(255), -- Email
                               CODE_EGRPOU:  "",  //VARCHAR(255), -- Код ОКПО
                               ADRESS:       "",  //VARCHAR(255), -- Юридический адрес организации 
                               ADRESS_FACT:  "",  //VARCHAR(255), -- Фактический адрес организации 
                               MIN_QNTY INT: 1,   //, -- Минимальное количество заказов
                               MIN_SUM INT:  1 ,  // Минимальная сумма заказа
                               MIN_DAY INT:  1,   // Минимальная срок поставки в днях
                               COMMENTS:    "",   // VARCHAR(255)) -- Коментарии 
                               });

//r.table("S_SUPPLIERS").delete()
