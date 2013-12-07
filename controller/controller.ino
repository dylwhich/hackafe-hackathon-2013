#define numPins 4
#define numMotors 2

const byte motor1pins[] = {2,3,4,5};
const byte motor2pins[] = {6,7,8,9};
const byte* motorpins[] = {motor1pins, motor2pins};

byte motorstep[] = {0, 0};

void setup(){
  for(byte motor = 0; motor < numMotors; motor++){
    for(byte pin = 0; pin < numPins; pin++){
      pinMode(motorpins[motor][pin], OUTPUT);
      
      if(motorstep[motor] == pin) digitalWrite(motorpins[motor][pin], HIGH);
      else digitalWrite(motorpins[motor][pin], LOW);
    }
  }
}

void loop(){
  stepmotor(1,200);
  stepmotor(1,-200);
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
    
    delay(5);
  }
}

//void stepmotor1(int numSteps){
//  for(int i = 0; i < abs(numSteps); i++){
//    if(numSteps > 0){
//      digitalWrite(motor1pins[motor1step], LOW);
//      motor1step += 1;
//      if(motor1step >= numPins) motor1step = 0;
//      digitalWrite(motor1pins[motor1step], HIGH);
//    }
//    else{
//      digitalWrite(motor1pins[motor1step], LOW);
//      motor1step -= 1;
//      if(motor1step >= numPins) motor1step = numPins - 1;
//      digitalWrite(motor1pins[motor1step], HIGH);
//    }
//    
//    delay(5);
//  }
//}
