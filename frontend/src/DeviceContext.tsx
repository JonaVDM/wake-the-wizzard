import { createContext, useEffect, useState } from 'react';
import { Device, getDevices } from './api';

interface Context {
  devices: Device[];
  addDevice: (device: Device) => void;
  removeDevice: (id: string) => void;
}

interface Props {
  children: any;
}

const defaultValue: Context = {
  devices: [],
  addDevice() {},
  removeDevice() {},
};

export const DeviceContext = createContext(defaultValue);

export function DeviceContextProvider({ children }: Props) {
  const [devices, setDevices] = useState<Device[]>(defaultValue.devices);

  const addDevice = (device: Device) => {
    setDevices([...devices]);
  };

  const removeDevice = (id: String) => {
    setDevices(devices.filter((d) => d.id != id));
  };

  useEffect(() => {
    loadDevices();
  }, []);

  const loadDevices = async () => {
    setDevices(await getDevices());
  };

  const provider: Context = {
    devices,
    addDevice,
    removeDevice,
  };

  return (
    <DeviceContext.Provider value={provider}>{children}</DeviceContext.Provider>
  );
}
