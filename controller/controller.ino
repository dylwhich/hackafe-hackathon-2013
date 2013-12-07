// number of pins to control a motor.
#define numPins 4
// number of motors available
#define numMotors 2

#define degPerStep 1.8
#define distPerDeg 1

// the pins for controlling each motor
const byte motorpins[][] = {{2,3,4,5},{6,7,8,9}};

// current step positions of each motor.
byte currentstep[] = {0, 0};

void setup(){
  for(byte i = 0; i < numMotors; i++){
    for(byte k = 0; i < numPins; k++){
      pinMode(motorpins[i][k], OUTPUT);
      
      if(currentstep[i] == k){
        digitalWrite(motorpins[i][k], HIGH);
      }
      else { 
        digitalWrite(motorpins[i][k], LOW);
      }
    }
  }
}

void loop(){
 
}
