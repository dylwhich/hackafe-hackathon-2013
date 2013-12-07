// number of pins to control a motor.
#define pincount 4
// number of motors available
#define numMotors 2

// the pins for controlling each motor
const byte motorpins[][] = {{2,3,4,5},{6,7,8,9}};

// current step positions of each motor.
byte currentstep[] = {0, 0};

void setup(){
  
}

void loop(){
  
}

void stepMotor(int motornumber, int steps) {
 // We want to move forward if the number is positive
 // we want to move backward if the number is negative
 if(steps >= pincount) {
   digitalWrite(motorpins[motornumber][currentpins[motornumber]],LOW]
   currentstep[motornumber] += 1;
   currentStep[motornumber] %= pincount;
   digitalWrite(motoirpins[motornumber][currentpins[motornumber]],HIGH]
 
 } else if(steps <= pincount) {
   digialWrite(motorPins[motornumber][currentpins[motornumber]],LOW]
   currentstep[motornumber] -= 1;
   currentstep[motornumber] %= pincount;
   digialWrite(motorPins[motornumber][currentpins[motornumber]],HIGH]
 }
}
