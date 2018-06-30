#include "gtest/gtest.h"

#include "../foo.h"
using namespace std;

TEST(AddTest, OneAndTwo){
    EXPECT_EQ(3, myadd(1, 2));

}

TEST(AddTest, TwoAndTwo){
    EXPECT_EQ(4, myadd(2, 2));

}

int main(int argc, char **argv) {
    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}