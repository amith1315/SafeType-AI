import {
    Table,
    TableBody,
    TableCaption,
    TableCell,
    TableFooter,
    TableHead,
    TableHeader,
    TableRow,
} from "@/components/ui/table"
import React from "react"

const TableView = ({ data }) => {
    // let sum = 0;
    // for(let i = 0; i < data.length; i++) {
    //     sum += data[i].mental_state
    // }
    return (
        // <div className="flex flex-row justify-evenly align-middle">
            <Table className="ml-8 w-11/12">
                <TableHeader>
                    <TableRow>
                        <TableHead className="w-[100px]">Sensitive Information</TableHead>
                        <TableHead className="text-left">NSFW Content</TableHead>
                        <TableHead className="text-left">Mental State</TableHead>
                        <TableHead className="text-left">Data and Time</TableHead>
                    </TableRow>
                </TableHeader>
                <TableBody>
                    {data.map((d) => (
                    <TableRow key={d.id}>

                        <TableCell className="text-left">{d.content.substring(0, 20)}</TableCell>
                        <TableCell className="text-left">{(d.nsfw_content ? "" : "null") || d.nsfw_content}</TableCell>
                        <TableCell className="text-left">{d.mental_state}</TableCell>
                        <TableCell className="text-left">{d.cur_data_time.substring(0, 10)} | {d.cur_data_time.substring(11, 19)}</TableCell>
                    </TableRow>
                    ))}
                </TableBody>
            </Table>
        // </div> //
    )
}

export default TableView

