import axios from './client';

const API_PATH = '/api/v1/akusaja';

export const getAllAkusajas = async (params) => {
    const response = await axios.get(API_PATH, { params });
    return response.data;
};

export const getAkusajaById = async (id) => {
    const response = await axios.get(`${API_PATH}/${id}`);
    return response.data;
};

export const createAkusaja = async (data) => {
    const response = await axios.post(API_PATH, data);
    return response.data;
};

export const updateAkusaja = async (id, data) => {
    const response = await axios.put(`${API_PATH}/${id}`, data);
    return response.data;
};

export const deleteAkusaja = async (id) => {
    const response = await axios.delete(`${API_PATH}/${id}`);
    return response.data;
};
