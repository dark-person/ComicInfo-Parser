import {
    type ColumnDef,
    flexRender,
    getCoreRowModel,
    getFilteredRowModel,
    getPaginationRowModel,
    useReactTable,
} from "@tanstack/react-table";
import { useState } from "react";
import { Button, Form, Table } from "react-bootstrap";

interface DataTableProps<DataT> {
    /** Column defintion for data table. */
    columns: ColumnDef<DataT>[];
    /** Data for past inputted value in database. */
    data: DataT[];
}

/** Data table for autofill value, which is values that inputted once. */
export default function DataTable<DataT>({ columns, data }: Readonly<DataTableProps<DataT>>) {
    const [globalFilter, setGlobalFilter] = useState("");

    const table = useReactTable({
        data,
        columns,
        getCoreRowModel: getCoreRowModel(),
        getPaginationRowModel: getPaginationRowModel(),
        getFilteredRowModel: getFilteredRowModel(),
        globalFilterFn: "includesString",
        onGlobalFilterChange: setGlobalFilter,
        state: { globalFilter },
    });

    return (
        <>
            <div className="pb-4">
                <Form.Control
                    value={table.getState().globalFilter}
                    onChange={(e) => table.setGlobalFilter(String(e.target.value))}
                    placeholder="Search..."
                />
            </div>

            <Table striped bordered hover variant="dark">
                <thead>
                    {table.getHeaderGroups().map((headerGroup) => (
                        <tr key={headerGroup.id}>
                            {headerGroup.headers.map((header) => {
                                return (
                                    <th key={header.id}>
                                        {header.isPlaceholder
                                            ? null
                                            : flexRender(header.column.columnDef.header, header.getContext())}
                                    </th>
                                );
                            })}
                        </tr>
                    ))}
                </thead>
                <tbody>
                    {table.getRowModel().rows?.length ? (
                        table.getRowModel().rows.map((row) => (
                            <tr key={row.id} data-state={row.getIsSelected() && "selected"}>
                                {row.getVisibleCells().map((cell) => (
                                    <td key={cell.id}>{flexRender(cell.column.columnDef.cell, cell.getContext())}</td>
                                ))}
                            </tr>
                        ))
                    ) : (
                        <tr>
                            <td colSpan={columns.length} className="h-24 text-center">
                                No results.
                            </td>
                        </tr>
                    )}
                </tbody>
            </Table>

            <div className="my-2 d-flex justify-content-center">
                <Button
                    variant="outline-light"
                    size="sm"
                    className="mx-2"
                    onClick={() => table.previousPage()}
                    disabled={!table.getCanPreviousPage()}>
                    {"<"}
                </Button>

                <div className="align-items-center m-1">
                    {table.getState().pagination.pageIndex + 1}/{table.getPageCount()}
                </div>

                <Button
                    variant="outline-light"
                    size="sm"
                    className="mx-2"
                    onClick={() => table.nextPage()}
                    disabled={!table.getCanNextPage()}>
                    {">"}
                </Button>
            </div>
        </>
    );
}
