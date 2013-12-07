#define numPins 4
#define numMotors 2
#define stepDelay 10
#define sign(X) (X < 0 ? -1 : X > 0 ? 1 : 0)

#define turnMotor 't'
#define interlace 'i'
#define raisePen 'u'
#define lowerPen 'd'

#define penDown 0
#define penUp 65

#include <Servo.h>

const byte motor1pins[] = {2,3,4,5};
const byte motor2pins[] = {6,7,8,9};
const byte penMotor = 11;
const byte* motorpins[] = {motor1pins, motor2pins};

byte motorstep[] = {0, 0};

Servo penServo;

void setup(){
  Serial.begin(9600);
  penServo.attach(penMotor);
  penServo.write(penUp);
  for(byte motor = 0; motor < numMotors; motor++){
    for(byte pin = 0; pin < numPins; pin++){
      pinMode(motorpins[motor][pin], OUTPUT);
      digitalWrite(motorpins[motor][pin], LOW);
    }
  }
}

void loop(){
  
  if(true && Serial.available()){
     char marker = Serial.read();
     
     if(marker == raisePen){
       for (int pos = penDown; pos < penUp; pos += 1){
         penServo.write(pos);
         delay(10);
       }
     }
     if(marker == lowerPen) {
       for (int pos = penUp; pos > penDown; pos += -1){
         penServo.write(pos);
         delay(10);
       }
     }
     if(marker == turnMotor){
       // wait for 3 bytes: 1 for a motor number, 2 for a number of turns
       while(Serial.available() < 3);
       
       byte motorNum = Serial.read();
       
       byte nums[2];
       nums[0] = Serial.read();
       nums[1] = Serial.read();
             
       int numSteps = *((int*)nums);
       
       stepmotor(motorNum, numSteps);
     }
     if(marker == interlace){
       while(Serial.available() < 4);
       
       byte nums[4];
       nums[0] = Serial.read();
       nums[1] = Serial.read();
       nums[2] = Serial.read();
       nums[3] = Serial.read();
       
       int motor1dist = *((int*)nums);
       int motor2dist = *((int*)(nums+2));
       
       signed char motor1dir = sign(motor1dist);
       signed char motor2dir = sign(motor2dist);
       
       motor1dist = abs(motor1dist);
       motor2dist = abs(motor2dist);
       
       if(motor1dist > motor2dist){         
         int onepertwo = motor1dist/motor2dist;
         
         while(motor1dist > 0 || motor2dist > 0){
           stepmotor(0, min(motor1dist, onepertwo) * motor1dir);
           stepmotor(1, min(1, motor2dist) * motor2dir);
           
           motor1dist = max(motor1dist - onepertwo, 0);
           motor2dist = max(motor2dist - 1, 0);
         }
       }
       else{
         int twoperone = motor2dist/motor1dist;
         
         while(motor1dist > 0 || motor2dist > 0){
           stepmotor(1, min(motor2dist, twoperone) * motor2dir);
           stepmotor(0, min(1, motor1dist) * motor1dir);
           
           motor2dist = max(motor2dist - twoperone, 0);
           motor1dist = max(motor1dist - 1, 0);
         }
       }
     }
  for(byte motor = 0; motor < numMotors; motor++){
    for(byte pin = 0; pin < numPins; pin++){
      pinMode(motorpins[motor][pin], OUTPUT);
      digitalWrite(motorpins[motor][pin], LOW);
    }
  }
  }
}

void stepmotor(byte motornum, int numSteps){
  for(int i = 0; i < abs(numSteps); i++){
    if(numSteps > 0){
      digitalWrite(motorpins[motornum][motorstep[motornum]], LOW);
      motorstep[motornum] += 1;
      if(motorstep[motornum] >= numPins) motorstep[motornum] = 0;
      digitalWrite(motorpins[motornum][motorstep[motornum]], HIGH);
    }
    else{
      digitalWrite(motorpins[motornum][motorstep[motornum]], LOW);
      motorstep[motornum] -= 1;
      if(motorstep[motornum] >= numPins) motorstep[motornum] = numPins - 1;
      digitalWrite(motorpins[motornum][motorstep[motornum]], HIGH);
    }
    
    delay(stepDelay);
  }
}
