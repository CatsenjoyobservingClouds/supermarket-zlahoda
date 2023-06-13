import React from 'react';
import DatabaseComponent from '../components/DatabaseComponent';

export default function Categories() {
    const rowData = {};
    const columnNames = ['№', 'Name', 'Average Product Price'];
    const tableName= "Category";
    const endpoint = "";
    return (
        <main>
            <DatabaseComponent
                endpoint={endpoint}
                handleData={() => {}}
                columnNames={columnNames}
                tableName={tableName} />
        </main>
    )
}