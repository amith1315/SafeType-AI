import React, { useEffect, useState } from 'react';
// import TableView from './TableView';
import { AppSidebar } from '@/components/app-sidebar';
import { SectionCards } from "@/components/section-cards"
import { SiteHeader } from "@/components/site-header"
import { SidebarInset, SidebarProvider } from "@/components/ui/sidebar"
import TableView from './TableView';

const MainPage = () => {

    const [data, setData] = useState<any[]>([]);

    useEffect(() => {
            const fetchData = async () => {
                const user = "vincent";
                const response = await fetch(`http://localhost:8000/getdata?name=${user}`, {
                    method: "GET",
                    headers: {
                        'Content-Type': 'application/json',
                    },
                });
                const res = await response.json();
                setData(res);
            };
            fetchData();
    }, []);

    return (
        <SidebarProvider className=''>
            <AppSidebar variant="inset" />
            <SidebarInset>
                <SiteHeader />
                <div className="flex flex-1 flex-col">
                <div className="@container/main flex flex-1 flex-col gap-2">
                    <div className="flex flex-col gap-4 py-4 md:gap-6 md:py-6">
                    <SectionCards data={data}/>
                    <div className="px-4 lg:px-6">
                    </div>
                    <h1 className="mt-8 text-left ml-10 text-2xl font-medium">User Logs</h1>
                    <div className="flex flex-col">
                        <TableView data={data}/>
                    </div>
                    </div>
                </div>
                </div>
            </SidebarInset>
        </SidebarProvider>
    )
}

export default MainPage