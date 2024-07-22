import {Howl} from "howler";
import audioKick from "./sounds/kicking.wav";
import audioNewPlayer from "./sounds/new-player.wav";
import audioLostPlayer from "./sounds/player-droped.mp3";
import audioRefereeStart from "./sounds/referee-whistle.mp3";
import audioBackground from "./sounds/gb.wav";
import audioPublic from "./sounds/public.mp3";
import debugPressed from "./sounds/on-debug-pressed.wav";
import goal from "./sounds/goal1.wav";
import audioConnectionLost from "./sounds/connection-lost.mp3";
import audioReconnected from "./sounds/reconnected.mp3";


class AudioManager {
  constructor() {
    this.context = new AudioContext();
    // this.playabble = false;
    this.audio = null;
    this.ambience_on = false
    this.ambient_audio_num =null
  }

  __canPlay() {
    if(this.context.state !== "suspended") {
      if(this.audio == null) {
        this.__initializaAutio()
      }
      return true;
    }
    this.context = new AudioContext();
    console.log(this.context.state)
    if(this.context.state === "suspended") {
       return false;
    }

    if(this.audio == null) {
      this.__initializaAutio()
    }
    return true;
  }

  __initializaAutio() {
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
      debugPressed: new Howl({
        src: [debugPressed]
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
    // removing ambient sound temporariay
    // this.ambient_audio_num = this.audio.background.play()
    // this.audio.background.fade(0, 0.06, 5000);
    this.ambience_on = true
  }

  stopAmbienceSound() {
    if(this.ambience_on) {
      this.audio.background.fade(0.06, 0, 1000);
      this.ambience_on = false
    }
  }

  isAmbienceOn() {
    return this.ambience_on
  }

  onGameRestart() {
    if(this.__canPlay()) {
      // console.log(`onGameRestart`)
      this.audio.refereeStart.play()
      this.onGameResume()
    }

  }

  onGameResume() {
    // console.log(`onGameResume`)
    this._startAmbienceSound()
  }


  onKick() {
    if(this.__canPlay()) {
      this.audio.kick.play()
    }

  }

  onNewPlayer() {
    if(this.__canPlay()) {


      this.audio.newPlayer.play()
    }
  }

  onLostPlayer() {
    if(this.__canPlay()) {


      this.audio.lostPlayer.play()
    }
  }

  onDebugPressed() {
    if(this.__canPlay()) {


      this.audio.debugPressed.play()
    }
  }

  onGoal() {
    if(this.__canPlay()) {


    const playNum = this.audio.goal.play("goal")
      setTimeout(() => {
        this.audio.goal.fade(0.3, 0, 5000, playNum);
      }, 3000)
    }

  }

  onGameOver() {
    this.stopAmbienceSound()
  }

  onBackendConnectionLost() {
    this.onUpstreamConnectionLost()
  }
  onBackendReconnected() {
    this.onUpstreamReconnected()
  }
  onUpstreamConnectionLost() {
    if(!this.__canPlay()) {
      return
    }
    this.stopAmbienceSound()
    try {
      this.audio.connectionLost.play()
    }catch (e) {
      console.log("GOT IS")
    }

  }
  onUpstreamReconnected() {
    if(!this.__canPlay()) {
      return
    }
    this.audio.reconnected.play()
  }
}

const manager = new AudioManager()

export default manager
