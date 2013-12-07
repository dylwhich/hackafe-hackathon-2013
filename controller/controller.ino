#define numPins 4
#define numMotors 2
#define stepDelay 10
#define sign(X) (X < 0 ? -1 : X > 0 1 : 0)

#define turnMotor 't'
#define interlace 'i'

const byte motor1pins[] = {2,3,4,5};
const byte motor2pins[] = {6,7,8,9};
const byte* motorpins[] = {motor1pins, motor2pins};

byte motorstep[] = {0, 0};

void setup(){
  Serial.begin(9600);
  
  for(byte motor = 0; motor < numMotors; motor++){
    for(byte pin = 0; pin < numPins; pin++){
      pinMode(motorpins[motor][pin], OUTPUT);
      
      if(motorstep[motor] == pin) digitalWrite(motorpins[motor][pin], HIGH);
      else digitalWrite(motorpins[motor][pin], LOW);
    }
  }
}

void loop(){
  if(Serial.available()){
     char marker = Serial.read();
     
     if(marker == turnMotor){
       // wait for 3 bytes: 1 for a motor number, 2 for a number of turns
       while(Serial.available() < 3);
       
       byte motorNum = Serial.read();
       
       byte nums[2];
       Serial.readBytes(nums, 2);
       
       int numSteps = *nums;
       
       stepmotor(motorNum, numSteps);
     }
     if(marker == interlace){
       while(Serial.available() < 4);
       
       byte nums[4];
       Serial.readBytes(nums, 4);
       
       int motor1dist = *nums;
       int motor2dist = *(nums + 2);
       
       byte motor1dir = sign(motor1dist);
       byte motor2dir = sign(motor2dist);
       
       motor1dist = abs(motor1dist);
       motor2dist = abs(motor2dist);
       
       if(motor1dist > motor2dist){         
         onepertwo = motor1dist/motor2dist;
         
         while(motor1dist > 0 || motor2dist > 0){
           stepmotor(0, min(motor1dist, onepertwo) * motor1dir;
           stepmotor(1, min(1, motor2dist) * motor2dir;
           
           motor1dist = max(motor1dist - onepertwo, 0);
           motor2dist = max(motor2dist - 1, 0);
         }
       }
       else{
         twoperone = motor2dist/motor1dist;
         
         while(motor1dist > 0 || motor2dist > 0){
           stepmotor(1, min(motor2dist, twoperone) * motor2dir;
           stepmotor(0, min(1, motor1dist) * motor1dir;
           
           motor2dist = max(motor2dist - twoperone, 0);
           motor1dist = max(motor1dist - 1, 0);
         }
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
