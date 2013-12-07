#define numPins 4
#define numMotors 2

const byte motor1pins[] = {2,3,4,5};
const byte motor2pins[] = {6,7,8,9};
const byte* motorPins[] = {motor1pins, motor2pins};

byte motorstep[] = {0, 0};

void setup(){
  for(byte motor = 0; motor < numMotors; motor++){
    for(byte pin = 0; pin < numPins; pin++){
      pinMode(motorPins[motor][pin], OUTPUT);
    }
  }
}

void loop(){
  stepmotor1(200);
  stepmotor1(-200);
}

void stepmotor1(int numSteps){
  for(int i = 0; i < abs(numSteps); i++){
    if(numSteps > 0){
      digitalWrite(motor1pins[motor1step], LOW);
      motor1step += 1;
      if(motor1step >= numPins) motor1step = 0;
      digitalWrite(motor1pins[motor1step], HIGH);
    }
    else{
      digitalWrite(motor1pins[motor1step], LOW);
      motor1step -= 1;
      if(motor1step >= numPins) motor1step = numPins - 1;
      digitalWrite(motor1pins[motor1step], HIGH);
    }
    
    delay(5);
  }
}

void stepmotor2(int numSteps){
  for(int i = 0; i < abs(numSteps); i++){
    if(numSteps > 0){
      digitalWrite(motor2pins[motor2step], LOW);
      motor2step += 1;
      if(motor2step >= numPins) motor2step = 0;
      digitalWrite(motor2pins[motor2step], HIGH);
    }
    else{
      digitalWrite(motor2pins[motor2step], LOW);
      motor2step -= 1;
      if(motor2step >= numPins) motor2step = numPins - 1;
      digitalWrite(motor2pins[motor2step], HIGH);
    }
    
    delay(5);
  }
}
