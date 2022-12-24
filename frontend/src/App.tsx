import { useContext, useState } from 'react';
import { AddDevice } from './components/AddDevice';
import { DeviceCard } from './components/DeviceCard';
import { DeviceContext } from './DeviceContext';

function App() {
  const { devices } = useContext(DeviceContext);
  const [toggleAdd, setToggleAdd] = useState<boolean>(false);

  return (
    <div className='w-full max-w-2xl mx-2 my-4'>
      <div className='flex justify-between'>
        <h1 className='text-3xl font-bold text-gray-300'>Wake On Lan</h1>

        <button
          className='bg-blue-600 text-white px-5 rounded-md hover:bg-blue-700 transition-all'
          onClick={() => setToggleAdd(!toggleAdd)}
        >
          Add
        </button>
      </div>

      {toggleAdd && <AddDevice />}
      {devices.map((device) => (
        <DeviceCard
          key={device.id}
          name={device.name}
          id={device.id}
          mac={device.mac}
        />
      ))}
    </div>
  );
}

export default App;
