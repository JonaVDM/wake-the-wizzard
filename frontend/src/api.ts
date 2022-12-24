export interface Device {
  id: string,
  name: string,
  mac: string,
}

export async function getDevices(): Promise<Device[]> {
  const req = await fetch('/api/pc');
  const body = await req.text();

  let devices: Device[] = [];

  for (let dev of body.split('\n')) {
    const [id, name, mac] = dev.split(';');
    if (!name) {
      continue;
    }
    devices.push({ id, name, mac })
  }

  return devices;
}

export async function addDevice(name: string, mac: string) {
  fetch('/api/pc', { method: 'POST', body: `${name};${mac}` });
}

export async function deleteDevice(id: string) {
  fetch(`/api/pc/${id}`, { method: 'DELETE' });
}

export async function wakeDevice(id: string) {
  fetch(`/api/wake/${id}`);
}
