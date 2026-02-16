import { useEffect, useState } from 'react';
import { useNavigate, Link } from 'react-router-dom';
import Button from '../components/Button';
import Card from '../components/Card';

const Dashboard = () => {
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
        <div className="min-h-screen bg-surface">
            {/* Navigation */}
            <nav className="bg-white shadow-md3-1">
                <div className="container mx-auto px-6 py-4">
                    <div className="flex justify-between items-center">
                        <h1 className="text-2xl font-bold text-primary-500">Go Starter</h1>
                        <div className="flex items-center gap-4">
                            <span className="text-gray-700">Welcome, {user.name}</span>
                            <Button variant="outline" onClick={handleLogout}>
                                Logout
                            </Button>
                        </div>
                    </div>
                </div>
            </nav>

            {/* Main Content */}
            <div className="container mx-auto px-6 py-12">
                <div className="max-w-4xl mx-auto">
                    <h2 className="text-4xl font-bold text-gray-900 mb-8">Dashboard</h2>

                    {/* User Info Card */}
                    <Card className="mb-8">
                        <h3 className="text-2xl font-semibold text-gray-900 mb-4">User Information</h3>
                        <div className="space-y-3">
                            <div className="flex">
                                <span className="font-medium text-gray-600 w-32">Name:</span>
                                <span className="text-gray-900">{user.name}</span>
                            </div>
                            <div className="flex">
                                <span className="font-medium text-gray-600 w-32">Email:</span>
                                <span className="text-gray-900">{user.email}</span>
                            </div>
                            <div className="flex">
                                <span className="font-medium text-gray-600 w-32">Role ID:</span>
                                <span className="text-gray-900">{user.role_id}</span>
                            </div>
                        </div>
                    </Card>

                    {/* Quick Actions */}
                    <Card>
                        <h3 className="text-2xl font-semibold text-gray-900 mb-4">Quick Actions</h3>
                        <div className="grid md:grid-cols-3 gap-4">
                            <Link to="/admin/users">
                                <Button fullWidth>Manage Users</Button>
                            </Link>
                            <Link to="/admin/roles">
                                <Button variant="secondary" fullWidth>Manage Roles</Button>
                            </Link>
                            <Link to="/admin/permissions">
                                <Button variant="outline" fullWidth>Manage Permissions</Button>
                            </Link>
                        </div>
                    </Card>

                    {/* Back to Home */}
                    <div className="mt-8 text-center">
                        <Link to="/" className="text-primary-500 hover:text-primary-600">
                            ← Back to Home
                        </Link>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default Dashboard;
