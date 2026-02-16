import apiClient from './client';

export const getUsers = async (page = 1, limit = 10) => {
    const response = await apiClient.get(`/users?page=${page}&limit=${limit}`);
    return response.data;
};

export const getRoles = async (page = 1, limit = 10) => {
    const response = await apiClient.get(`/roles?page=${page}&limit=${limit}`);
    return response.data;
};

export const getPermissions = async (page = 1, limit = 10) => {
    const response = await apiClient.get(`/permissions?page=${page}&limit=${limit}`);
    return response.data;
};
