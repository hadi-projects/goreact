/**
 * usePermission — returns a function to check if the current user has a given permission.
 * Permissions are stored in localStorage as part of the user object.
 *
 * Usage:
 *   const can = usePermission();
 *   can('create-produk') // true/false
 */
const usePermission = () => {
    const userData = localStorage.getItem('user');
    const permissions = userData ? (JSON.parse(userData).permissions || []) : [];
    return (permission) => permissions.includes(permission);
};

export default usePermission;
