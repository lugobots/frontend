import {Howl} from "howler";
import audioKick from "./sounds/kick.mp3";
import audioNewPlayer from "./sounds/new-player.wav";
import audioLostPlayer from "./sounds/player-droped.mp3";
import audioRefereeStart from "./sounds/referee-whistle.mp3";
import audioBackground from "./sounds/gb.wav";
import audioPublic from "./sounds/public.mp3";
import goal from "./sounds/goal1.wav";
import audioConnectionLost from "./sounds/connection-lost.mp3";
import audioReconnected from "./sounds/reconnected.mp3";


class AudioManager {
  constructor() {
    this.ambience_on = false
    this.ambient_audio_num =null
    this.audio = {
      background: new Howl({
        src: [audioBackground],
        volume: 1,
        loop: true,
      }),
      audioPublic: new Howl({
        src: [audioPublic],
        // volume: 0.5,
        sprite: {
          bg: [200, 400],
          claps_good_try: [145, 20],
        }
      }),

      kick: new Howl({
        src: [audioKick],
        volume: 1,

      }),
      newPlayer: new Howl({
        src: [audioNewPlayer]
      }),
      lostPlayer: new Howl({
        src: [audioLostPlayer]
      }),
      refereeStart: new Howl({
        src: [audioRefereeStart]
      }),
      connectionLost: new Howl({
        src: [audioConnectionLost]
      }),
      reconnected: new Howl({
        src: [audioReconnected]
      }),

      goal: new Howl({
        src: [goal],
        volume: 0.3,
        sprite: {
          goal: [4200, 8000],
        }
      })
    }
  }

  _startAmbienceSound() {
    this.ambient_audio_num = this.audio.background.play()
    this.audio.background.fade(0, 0.08, 5000);
    this.ambience_on = true
  }

  _stopAmbienceSound() {
    if(this.ambience_on) {
      this.audio.background.fade(0.08, 0, 1000);
      this.ambience_on = false
    }
  }

  isAmbienceOn() {
    return this.ambience_on
  }

  onGameRestart() {
    // console.log(`onGameRestart`)
    this.audio.refereeStart.play()
    this.onGameResume()
  }

  // onGameStarts() {
  //   console.log(`onGameStartss`)
  //
  // }

  onGameResume() {
    // console.log(`onGameResume`)
    this._startAmbienceSound()
  }


  onKick() {
    this.audio.kick.play()
  }

  onNewPlayer() {
    this.audio.newPlayer.play()
  }

  onLostPlayer() {
    this.audio.lostPlayer.play()
  }

  onGoal() {
    const playNum = this.audio.goal.play("goal")
    setTimeout(() => {
      this.audio.goal.fade(0.3, 0, 5000, playNum);
    }, 3000)
  }


  onBackendConnectionLost() {
    this.onUpstreamConnectionLost()
  }
  onBackendReconnected() {
    this.onUpstreamReconnected()
  }
  onUpstreamConnectionLost() {
    this._stopAmbienceSound()
    this.audio.connectionLost.play()
  }
  onUpstreamReconnected() {
    this.audio.reconnected.play()
  }
}

const manager = new AudioManager()

export default manager
