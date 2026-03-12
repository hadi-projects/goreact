import axios from './client';

const API_PATH = '/mina';

export const getAllMinas = async (params) => {
    const response = await axios.get(API_PATH, { params });
    return response.data;
};

export const getMinaById = async (id) => {
    const response = await axios.get(`${API_PATH}/${id}`);
    return response.data;
};

export const createMina = async (data) => {
    const response = await axios.post(API_PATH, data);
    return response.data;
};

export const updateMina = async (id, data) => {
    const response = await axios.put(`${API_PATH}/${id}`, data);
    return response.data;
};

export const deleteMina = async (id) => {
    const response = await axios.delete(`${API_PATH}/${id}`);
    return response.data;
};

export const exportMina = async (format = 'excel') => {
    return axios.get(`${API_PATH}/export?format=${format}`, {
        responseType: 'blob',
    });
};
