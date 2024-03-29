import React from 'react';
import DatabaseComponent from '../components/DatabaseComponent';

export default function CustomerCards() {
    const columnNames = ['Card Number', 'Full Name', 'Phone Number', 'City', 'Street', 'Zip Code', 'Discount Percent'];
    const columnNamesChange = ['Name', 'Surname', 'Patronymic', 'Phone Number', 'City', 'Street', 'Zip Code', 'Discount Percent'];

    const tableName= "Customer";
    const endpoint = "http://localhost:8080/" + localStorage.getItem("role")?.toLowerCase() + "/customerCard";

    const decodeData = (data: any[]) => {
        const chosenData = data.map((item) => ({
            "Id": item.card_number,
            'Card Number': item.card_number,
            'Full Name': item.cust_surname + " " + item.cust_name + " " + item.cust_patronymic["String"],
            'Name': item.cust_name,
            'Surname': item.cust_surname,
            'Patronymic': item.cust_patronymic["String"],
            'Phone Number': item.phone_number,
            'City': item.city["String"],
            'Street': item.street["String"],
            'Zip Code': item.zip_code["String"],
            'Discount Percent': item.percent + "%"
        }));
        return chosenData;
    }

    const encodeData = (data: any[]) => {
        const chosenData = data.map((item) => ({
            'card_number': item["Id"],
            'cust_surname': item["Surname"],
            'cust_name': item["Name"],
            'cust_patronymic': {
                "String" : (item["Patronymic"] ? item["Patronymic"] : ""),
                "Valid" : (item["Patronymic"] ? true : false)
            },
            'phone_number': item["Phone Number"],
            'city': {
                "String" : (item["City"] ? item["City"] : ""),
                "Valid" : (item["City"] ? true : false)
            },
            'street': {
                "String" : (item["Street"] ? item["Street"] : ""),
                "Valid" : (item["Street"] ? true : false)
            },
            'zip_code': {
                "String" : (item["Zip Code"] ? item["Zip Code"] : ""),
                "Valid" : (item["Zip Code"] ? true : false)
            },
            'percent': parseInt(item["Discount Percent"])
        }));
        return chosenData;
    }

    return (
        <main>
            <DatabaseComponent
                endpoint={endpoint}
                decodeData={decodeData}
                encodeData={encodeData}
                columnNames={columnNames}
                columnNamesChange={columnNamesChange}
                tableName={tableName} />
        </main>
    )
}