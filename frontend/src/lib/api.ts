import { writable } from 'svelte/store';

export const devices = writable<Device[]>([]);

export async function loadDevices() {
  const req = await fetch('/api/pc');
  const body: Device[] = await req.json();
  devices.set(body);
}

export async function addDevice(name: string, mac: string): Promise<Response> {
  return fetch('/api/pc', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({ name, mac })
  });
}

export async function deleteDevice(id: string) {
  fetch(`/api/pc/${id}`, { method: 'DELETE' });
  devices.update(devices => devices.filter(device => device.id != id));
}

export async function wakeDevice(id: string) {
  fetch(`/api/wake/${id}`);
}
