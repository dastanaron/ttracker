import { useEffect, useState } from 'react';

import './App.css';
import { convertSecondsToHumanTime } from './libraries/timer';
import { TaskModel } from './models/TaskModel';
import { GetTodayTask, Log, SaveTask } from '../wailsjs/go/main/App';

const DEFAULT_TASK_MODEL: TaskModel = {
  Id:        0,
  Name:      '',
  StartTime: new Date(),
  Duration:  0,
  Project:   '',
};

function App() {
  const [taskName, setTaskName] = useState('');
  const [duration, setDuration] = useState(0);
  const [isRunTimer, setRunTimer] = useState(false);
  const [intervalId, setIntervalId] = useState<number>();
  const [tasks, setTasks] = useState<TaskModel[]>([]);
  const [taskModel, setTaskModel] = useState<TaskModel>(DEFAULT_TASK_MODEL);

  const [requestCount, setRequestCount] = useState(1);

  const startTimer = () => {
    setRunTimer(true);
    let localDuration = 0;
    const intervalId = setInterval(() => {
      localDuration++;
      setDuration(localDuration);
    }, 1000);
    setIntervalId(intervalId);
  };

  const stopTimer = async () => {
    if (taskName.length < 2 || taskName.search(/:/) === -1) {
      alert('Task name is wrong, example:\nTASK-001: Task name from your tracker system');
      return;
    }

    const [projectName] = taskName.split(':');

    const cloneTaskModel = Object.assign({}, taskModel);
    cloneTaskModel.Name = taskName;
    cloneTaskModel.Duration = duration;
    cloneTaskModel.Project = projectName;

    try {
      await SaveTask(cloneTaskModel);
      clearInterval(intervalId);
      setIntervalId(undefined);
      setRunTimer(false);
      setTaskModel(DEFAULT_TASK_MODEL);
    } catch(e: any) {
      alert('Cannot save task');
      Log(e.toString());
    } finally {
      setRequestCount(requestCount+1);
    }
  };


  useEffect(() => {
    GetTodayTask().then(res => {
      setTasks(res);
    }); 
  }, [requestCount]);

  return (
    <div id="App">
      <div className="row">
        <div className='col'>
          <input
            style={{
              width: '200px',
            }}
            type="text" 
            value={taskName} 
            className="task-name-input" 
            placeholder='Input task name'
            onChange={(event) => {
              setTaskName(event.target.value);
            }} /> 
        </div>
        <div className='col'>
          <button type='button' className='start-btn' onClick={isRunTimer ? stopTimer : startTimer}>{isRunTimer ? 'Stop' : 'Start'}</button>
        </div>
      </div>
      <div className='timer'>
        {convertSecondsToHumanTime(duration)}
      </div>
      <div className='report'>
        <table className='report-time'>
          <thead>
            <tr>
              <td>name</td>
              <td>duration</td>
              <td>control</td>
            </tr>
          </thead>
          <tbody>
            {
              tasks.map((task) => (
                <tr>
                  <td>{task.Name}</td>
                  <td>{convertSecondsToHumanTime(task.Duration)}</td>
                  <td><button type='button' className='start-btn' onClick={async () => {
                    setTaskName(task.Name);
                    setTaskModel(task);
                    if (isRunTimer) {
                      await stopTimer();
                    }
                    startTimer();
                  }}>Start</button></td>
                </tr>
              ))
            }

          </tbody>
        </table>
      </div>
    </div>
  );
}

export default App;
