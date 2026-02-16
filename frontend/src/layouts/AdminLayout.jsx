import { useEffect, useState } from 'react';
import { useNavigate, Outlet } from 'react-router-dom';
import Sidebar from '../components/Sidebar';
import Button from '../components/Button';

const AdminLayout = () => {
    const navigate = useNavigate();
    const [user, setUser] = useState(null);

    useEffect(() => {
        const token = localStorage.getItem('token');
        const userData = localStorage.getItem('user');

        if (!token) {
            navigate('/login');
            return;
        }

        if (userData) {
            setUser(JSON.parse(userData));
        }
    }, [navigate]);

    const handleLogout = () => {
        localStorage.removeItem('token');
        localStorage.removeItem('user');
        navigate('/login');
    };

    if (!user) {
        return (
            <div className="min-h-screen bg-surface flex items-center justify-center">
                <div className="text-primary-500">Loading...</div>
            </div>
        );
    }

    return (
        <div className="flex min-h-screen bg-gray-50">
            {/* Sidebar */}
            <Sidebar />

            {/* Main Content */}
            <div className="flex-1 flex flex-col">
                {/* Header */}
                <header className="bg-white border-b border-outline-variant px-6 py-3">
                    <div className="flex justify-between items-center">
                        <h2 className="text-lg font-semibold text-gray-900">Administration</h2>
                        <div className="flex items-center gap-3">
                            <span className="text-sm text-gray-600">{user.email}</span>
                            <Button variant="outline" onClick={handleLogout}>
                                Logout
                            </Button>
                        </div>
                    </div>
                </header>

                {/* Page Content */}
                <main className="flex-1 p-6">
                    <Outlet />
                </main>
            </div>
        </div>
    );
};

export default AdminLayout;
