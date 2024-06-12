import i18n from "i18next";
import { initReactI18next } from "react-i18next";

const resources = {
    en: {
        translation: {
            "Warehouse": "Warehouse",
            "Cars": "Cars",
            "Managers": "Managers",
            "Create Warehouse": "Create Warehouse",
            "Create Slot": "Create Slot",
            "Create Car": "Create Car",
            "Email": "Email",
            "Submit": "Submit",
            "Go": "Go",
            "Delete": "Delete",
            "Car": "Car",
            "Create Item": "Create Item",
            "Name": "Name",
            "Description": "Description",
            "Weight": "Weight",
            "Slot": "Slot",
            "Create Manager": "Create Manager",
            "Max weight": "Max weight",
            "kg": "kg",
            "Create Worker": "Create Worker",
            "No Device": "No Device",
            "Create Device": "Create Device",
            "Device JWT": "Device JWT",
            "Item": "Item",
            "Empty": "Empty",
            "Update Item": "Update Item",
            "Update Slot": "Update Slot",
            "Update Car": "Update Car",
            "Update Manager": "Update Manager",
            "Update Worker": "Update Worker",
            "Update Device": "Update Device",
            "Update": "Update",
            "Manager": "Manager",
            "Worker": "Worker",
            "Last weight": "Last weight",
            "Free managers": "Free managers",
            "Workers": "Workers",
            "Device": "Device",
        }
    },
    ua: {
        translation: {
            "Warehouse": "Сховище",
            "Cars": "Машини",
            "Managers": "Менеджери",
            "Create Warehouse": "Створити сховище",
            "Create Slot": "Створити місце зберігання",
            "Create Car": "Створити машину",
            "Email": "Пошта",
            "Submit": "Підтвердити",
            "Go": "Перейти",
            "Delete": "Видалити",
            "Car": "Машина",
            "Create Item": "Стваорити товар",
            "Name": "Ім'я",
            "Description": "Опис",
            "Weight": "Вага",
            "Slot": "Місце зберігання",
            "Create Manager": "Створити менеджера",
            "Max weight": "Макс. вага",
            "kg": "кг",
            "Create Worker": "Створити працівника",
            "No Device": "Пристрій відсутній",
            "Create Device": "Створити пристрій",
            "Device JWT": "JWT пристрою",
            "Item": "Товар",
            "Empty": "Пусто",
            "Update Item": "Оновити товар",
            "Update Slot": "Оновити місце зберігання",
            "Update Car": "Оновити машину",
            "Update Manager": "Оновити менеджера",
            "Update Worker": "Оновити працівника",
            "Update Device": "Оновити пристрій",
            "Update": "Оновити",
            "Manager": "Менеджер",
            "Worker": "Працівник",
            "Last weight": "Остання вага",
            "Free managers": "Доступні менеджери",
            "Workers": "Працівники",
            "Device": "Пристрої",
            "Login": "Увійти",
            "Device Connected": "Пристрій підключено"
        }
    }
};

i18n
    .use(initReactI18next)
    .init({
        resources,
        lng: "en",

        interpolation: {
            escapeValue: false
        }
    });

export default i18n;