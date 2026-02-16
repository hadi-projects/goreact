import { useState } from 'react';
import { useQuery } from '@tanstack/react-query';
import Table from '../../components/Table';
import Pagination from '../../components/Pagination';
import Card from '../../components/Card';
import { getRoles } from '../../api/admin';

const Roles = () => {
    const [currentPage, setCurrentPage] = useState(1);
    const [itemsPerPage] = useState(10);

    const { data, isLoading, error } = useQuery({
        queryKey: ['roles', currentPage, itemsPerPage],
        queryFn: () => getRoles(currentPage, itemsPerPage),
    });

    const columns = [
        { header: 'ID', accessor: 'id' },
        { header: 'Name', accessor: 'name' },
        { header: 'Description', accessor: 'description' },
        {
            header: 'Created At',
            render: (row) => new Date(row.created_at).toLocaleDateString(),
        },
    ];

    if (error) {
        return (
            <div className="text-center py-12">
                <p className="text-red-500">Error loading roles: {error.message}</p>
            </div>
        );
    }

    const roles = data?.data || [];
    const meta = data?.meta?.pagination || { total_data: 0, total_pages: 1 };

    return (
        <div>
            <div className="mb-6">
                <h1 className="text-3xl font-bold text-gray-900">Roles Management</h1>
                <p className="text-gray-600 mt-2">Manage user roles and access levels</p>
            </div>

            <Card className="p-0 overflow-hidden">
                <Table columns={columns} data={roles} loading={isLoading} />
                {!isLoading && roles.length > 0 && (
                    <Pagination
                        currentPage={currentPage}
                        totalPages={meta.total_pages}
                        totalItems={meta.total_data}
                        itemsPerPage={itemsPerPage}
                        onPageChange={setCurrentPage}
                    />
                )}
            </Card>
        </div>
    );
};

export default Roles;
