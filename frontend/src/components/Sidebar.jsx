import { useState } from 'react';
import { Link, useLocation } from 'react-router-dom';

const Sidebar = () => {
    const location = useLocation();
    const [isAdminOpen, setIsAdminOpen] = useState(true);

    const isActive = (path) => location.pathname === path;
    const isAdminActive = () => location.pathname.startsWith('/admin/');

    return (
        <aside className="w-64 bg-gray-50 border-r border-gray-200 h-screen flex flex-col py-6">
            {/* Logo */}
            <div className="px-6 mb-8">
                <div className="flex items-center gap-3">
                    <div className="w-10 h-10 bg-primary-500 rounded-lg flex items-center justify-center">
                        <span className="text-white font-bold text-xl">G</span>
                    </div>
                    <h1 className="text-xl font-bold text-gray-900">Admin</h1>
                </div>
            </div>

            {/* Navigation */}
            <nav className="flex-1 overflow-y-auto">
                <ul className="space-y-1 px-3">
                    {/* Dashboard */}
                    <li>
                        <Link
                            to="/dashboard"
                            className={`flex items-center gap-3 px-4 py-3 rounded-xl transition-all ${isActive('/dashboard')
                                ? 'bg-gray-200 text-gray-900'
                                : 'text-gray-600 hover:bg-gray-100 hover:text-gray-900'
                                }`}
                        >
                            <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" />
                            </svg>
                            <span className="font-medium">Dashboard</span>
                        </Link>
                    </li>

                    {/* Administrator Parent */}
                    <li>
                        <button
                            onClick={() => setIsAdminOpen(!isAdminOpen)}
                            className={`w-full flex items-center justify-between px-4 py-3 rounded-xl transition-all ${isAdminActive()
                                ? 'bg-gray-200 text-gray-900'
                                : 'text-gray-600 hover:bg-gray-100 hover:text-gray-900'
                                }`}
                        >
                            <div className="flex items-center gap-3">
                                <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
                                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                                </svg>
                                <span className="font-medium">Administrator</span>
                            </div>
                            <svg
                                className={`w-4 h-4 transition-transform ${isAdminOpen ? 'rotate-180' : ''}`}
                                fill="none"
                                stroke="currentColor"
                                viewBox="0 0 24 24"
                            >
                                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 9l-7 7-7-7" />
                            </svg>
                        </button>

                        {/* Submenu */}
                        {isAdminOpen && (
                            <ul className="mt-2 ml-8 space-y-1">
                                <li>
                                    <Link
                                        to="/admin/users"
                                        className={`flex items-center gap-3 px-4 py-2 rounded-lg transition-all ${isActive('/admin/users')
                                            ? 'bg-gray-200 text-gray-900'
                                            : 'text-gray-500 hover:bg-gray-100 hover:text-gray-700'
                                            }`}
                                    >
                                        <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z" />
                                        </svg>
                                        <span className="text-sm">Users</span>
                                    </Link>
                                </li>
                                <li>
                                    <Link
                                        to="/admin/roles"
                                        className={`flex items-center gap-3 px-4 py-2 rounded-lg transition-all ${isActive('/admin/roles')
                                            ? 'bg-gray-200 text-gray-900'
                                            : 'text-gray-500 hover:bg-gray-100 hover:text-gray-700'
                                            }`}
                                    >
                                        <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />
                                        </svg>
                                        <span className="text-sm">Roles</span>
                                    </Link>
                                </li>
                                <li>
                                    <Link
                                        to="/admin/permissions"
                                        className={`flex items-center gap-3 px-4 py-2 rounded-lg transition-all ${isActive('/admin/permissions')
                                            ? 'bg-gray-200 text-gray-900'
                                            : 'text-gray-500 hover:bg-gray-100 hover:text-gray-700'
                                            }`}
                                    >
                                        <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 7a2 2 0 012 2m4 0a6 6 0 01-7.743 5.743L11 17H9v2H7v2H4a1 1 0 01-1-1v-2.586a1 1 0 01.293-.707l5.964-5.964A6 6 0 1121 9z" />
                                        </svg>
                                        <span className="text-sm">Permissions</span>
                                    </Link>
                                </li>
                            </ul>
                        )}
                    </li>
                </ul>
            </nav>

            {/* User Avatar */}
            <div className="px-6 mt-auto pt-4 border-t border-gray-200">
                <div className="flex items-center gap-3">
                    <div className="w-10 h-10 rounded-full bg-gray-300 flex items-center justify-center text-gray-600 font-medium">
                        A
                    </div>
                    <div className="flex-1 min-w-0">
                        <p className="text-sm font-medium text-gray-900 truncate">Admin</p>
                        <p className="text-xs text-gray-500">Administrator</p>
                    </div>
                </div>
            </div>
        </aside>
    );
};

export default Sidebar;
