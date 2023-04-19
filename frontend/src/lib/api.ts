import { writable } from 'svelte/store';

export const devices = writable<Device[]>([]);

export async function loadDevices() {
  const req = await fetch('/api/pc');
  const body: Device[] = await req.json();
  devices.set(body);
}

export async function addDevice(name: string, mac: string) {
  const req = await fetch('/api/pc', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({ name, mac })
  });

  const data = await req.json();
  devices.update(devices => [...devices, { name, mac, id: data.id }]);
}

export async function deleteDevice(id: string) {
  devices.update(devices => devices.filter(device => device.id != id));
  await fetch(`/api/pc/${id}`, { method: 'DELETE' });
}

export async function wakeDevice(id: string) {
  await fetch(`/api/wake/${id}`);
}
