import { Link, useLocation } from 'react-router-dom';
import PropTypes from 'prop-types';

const Sidebar = ({ userName }) => {
    const location = useLocation();

    const menuItems = [
        { path: '/admin/users', label: 'Users', icon: '👥' },
        { path: '/admin/roles', label: 'Roles', icon: '🔐' },
        { path: '/admin/permissions', label: 'Permissions', icon: '🔑' },
    ];

    const isActive = (path) => location.pathname === path;

    return (
        <div className="w-64 bg-primary-500 min-h-screen flex flex-col">
            {/* Logo/Brand */}
            <div className="p-6 border-b border-primary-600">
                <h1 className="text-2xl font-bold text-white">Go Starter</h1>
                <p className="text-sm text-primary-100 mt-1">Admin Panel</p>
            </div>

            {/* User Info */}
            <div className="p-6 border-b border-primary-600">
                <div className="flex items-center gap-3">
                    <div className="w-10 h-10 bg-primary-300 rounded-full flex items-center justify-center">
                        <span className="text-primary-700 font-semibold">
                            {userName?.charAt(0).toUpperCase()}
                        </span>
                    </div>
                    <div>
                        <p className="text-white font-medium text-sm">{userName}</p>
                        <p className="text-primary-100 text-xs">Administrator</p>
                    </div>
                </div>
            </div>

            {/* Navigation Menu */}
            <nav className="flex-1 p-4">
                <div className="space-y-2">
                    {menuItems.map((item) => (
                        <Link
                            key={item.path}
                            to={item.path}
                            className={`
                flex items-center gap-3 px-4 py-3 rounded-md3
                transition-all duration-200
                ${isActive(item.path)
                                    ? 'bg-primary-600 text-white shadow-md3-1'
                                    : 'text-primary-100 hover:bg-primary-600 hover:text-white'
                                }
              `}
                        >
                            <span className="text-xl">{item.icon}</span>
                            <span className="font-medium">{item.label}</span>
                        </Link>
                    ))}
                </div>
            </nav>

            {/* Footer */}
            <div className="p-4 border-t border-primary-600">
                <Link
                    to="/"
                    className="flex items-center gap-2 px-4 py-2 text-primary-100 hover:text-white transition-colors"
                >
                    <span>←</span>
                    <span className="text-sm">Back to Home</span>
                </Link>
            </div>
        </div>
    );
};

Sidebar.propTypes = {
    userName: PropTypes.string.isRequired,
};

export default Sidebar;
