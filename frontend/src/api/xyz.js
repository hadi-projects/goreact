import axios from './client';

const API_PATH = '/api/v1/xyz';

export const getAllXyzs = async (params) => {
    const response = await axios.get(API_PATH, { params });
    return response.data;
};

export const getXyzById = async (id) => {
    const response = await axios.get(`${API_PATH}/${id}`);
    return response.data;
};

export const createXyz = async (data) => {
    const response = await axios.post(API_PATH, data);
    return response.data;
};

export const updateXyz = async (id, data) => {
    const response = await axios.put(`${API_PATH}/${id}`, data);
    return response.data;
};

export const deleteXyz = async (id) => {
    const response = await axios.delete(`${API_PATH}/${id}`);
    return response.data;
};
