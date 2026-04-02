import { createContext, useContext, useState, useEffect } from 'react';
import PropTypes from 'prop-types';
import { getSettingsByCategory, getPublicSettings } from '../api/settings';

const SettingsContext = createContext();

export const SettingsProvider = ({ children }) => {
    const [settings, setSettings] = useState({
        app_name: '...',
        logo: null,
        favicon: null,
    });

    const refreshSettings = async () => {
        try {
            // Fetch website and storage settings in parallel
            const [webRes, storageRes] = await Promise.all([
                getPublicSettings('website'),
                getPublicSettings('storage').catch(() => ({ data: { data: [] } }))
            ]);

            const websiteSettings = webRes.data.data;
            const storageSettings = storageRes.data.data;

            const newSettings = { ...settings };
            websiteSettings.forEach(s => {
                if (s.key === 'app_name' && s.value) newSettings.app_name = s.value;
                if (s.key === 'app_logo' && s.value) newSettings.logo = s.value;
                if (s.key === 'app_favicon' && s.value) newSettings.favicon = s.value;
            });

            storageSettings.forEach(s => {
                if (s.key === 'storage_max_file_size_mb') {
                    newSettings.max_file_size_mb = parseInt(s.value, 10) || 50;
                }
            });

            setSettings(newSettings);

            // Update document title and favicon
            if (newSettings.app_name) {
                // Individual pages handle specific titles
            }

            if (newSettings.favicon) {
                const link = document.querySelector("link[rel*='icon']") || document.createElement('link');
                link.type = 'image/x-icon'; link.rel = 'shortcut icon';
                link.href = `${import.meta.env.VITE_API_URL}/public/share/${newSettings.favicon}`;
                document.getElementsByTagName('head')[0].appendChild(link);
            }
        } catch (err) {
            console.error('Failed to fetch settings:', err);
        }
    };

    useEffect(() => {
        refreshSettings();
    }, []);

    return (
        <SettingsContext.Provider value={{ ...settings, refreshSettings }}>
            {children}
        </SettingsContext.Provider>
    );
};

SettingsProvider.propTypes = {
    children: PropTypes.node.isRequired,
};

export const useSettings = () => {
    const context = useContext(SettingsContext);
    if (!context) {
        throw new Error('useSettings must be used within a SettingsProvider');
    }
    return context;
};
