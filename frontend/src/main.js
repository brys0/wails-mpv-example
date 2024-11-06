import './style.css';
import './app.css';

import logo from './assets/images/logo-universal.png';
import {StartMPV} from '../wailsjs/go/main/App';

document.querySelector('#app').innerHTML = `
    <img id="logo" class="logo">
      <button onclick="play()">Start MPV player</button>
    </div>
`;
document.getElementById('logo').src = logo;


// Setup the greet function
window.play = function (e) {
    // Call App.Greet(name)
    try {
        StartMPV(name)
            .then(() => {
                // Update result with data back from App.Greet()
                e.currentTarget.innerText = "Started";
            })
            .catch((err) => {
                console.error(err);
            });
    } catch (err) {
        console.error(err);
    }
};
